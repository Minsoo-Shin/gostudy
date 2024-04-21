package storage_cloud

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

type StorageCloud struct {
	S3Client *s3.Client
}

func New() StorageCloud {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
	}
	return StorageCloud{
		S3Client: s3.NewFromConfig(sdkConfig),
	}
}

func (sc StorageCloud) UploadLargeObject(bucketName string, objectKey string, largeObject []byte) error {
	// db에 저장
	largeBuffer := bytes.NewReader(largeObject)
	var partMiBs int64 = 10
	uploader := manager.NewUploader(sc.S3Client, func(u *manager.Uploader) {
		u.PartSize = partMiBs * 1024 * 1024
	})

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   largeBuffer,
	})
	if err != nil {
		log.Printf("Couldn't upload large object to %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}

	return err
}

func (sc StorageCloud) GetObject(bucketName string, objectKey string) {

}
