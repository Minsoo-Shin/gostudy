package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Minsoo-Shin/ms_drive/entity"
	"github.com/Minsoo-Shin/ms_drive/internal/drive/dto"
	"github.com/Minsoo-Shin/ms_drive/internal/drive/repository"
	"github.com/Minsoo-Shin/ms_drive/pkg/storage_cloud"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	serviceBucketName = "ms-drive"
)

type controller struct {
	storageCloud storage_cloud.StorageCloud
	repository   repository.DriverRepository
}

func New(r *gin.Engine, sc storage_cloud.StorageCloud, repository repository.DriverRepository) {
	ct := controller{
		storageCloud: sc,
		repository:   repository,
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/files/upload", ct.UploadFile)
	r.POST("/files/download", ct.DownloadFile)
}

func (ct controller) UploadFile(c *gin.Context) {
	// query: namespaceID
	// query: filepath
	req := dto.CreateFileRequest{}
	// using BindJson method to serialize body with struct
	if err := c.ShouldBind(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// get namespace name
	namespace := func(ctx context.Context, namespaceID int64) string {
		return "userA"
	}(c.Request.Context(), req.NamespaceID)

	// open file
	file, err := req.File.Open()
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fileBinary, _ := ioutil.ReadAll(file)

	s3ObjectKey, err := url.JoinPath(namespace, req.FilePath, req.File.Filename)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = ct.storageCloud.UploadLargeObject(serviceBucketName, s3ObjectKey, fileBinary)
	if err != nil {
		log.Fatalf("upload object err: %v", err)
	}

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println("해시를 계산하는 중 오류가 발생했습니다:", err)
		return
	}

	checksum := hex.EncodeToString(hash.Sum(nil))

	newFile, err := ct.repository.Create(&entity.File{
		FileName:      req.File.Filename,
		RelativePath:  req.FilePath,
		IsDirectory:   true,
		LatestVersion: 1,
		Checksum:      checksum,
		NamespaceID:   req.NamespaceID,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	})
	if err != nil {
		log.Fatalf("repository create method err: %v", err)
	}
	c.JSON(http.StatusOK, newFile)
}

func (ct controller) DownloadFile(c *gin.Context) {
}
