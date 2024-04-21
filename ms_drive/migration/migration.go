package migration

import (
	"fmt"
	"github.com/Minsoo-Shin/ms_drive/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func NewMigration(db *gorm.DB) {
	CheckError(db.AutoMigrate(entity.Device{}))
	CheckError(db.AutoMigrate(entity.User{}))
	CheckError(db.AutoMigrate(entity.File{}))
	CheckError(db.AutoMigrate(entity.FileVersion{}))
	CheckError(db.AutoMigrate(entity.Namespace{}))
	CheckError(db.AutoMigrate(entity.Block{}))
}

func Sample(db *gorm.DB) {
	deviceID := uuid.New().String()
	// sample data
	CheckError(db.Create(&entity.User{
		ID:       1,
		Username: "ms",
		Namespaces: []entity.Namespace{
			{
				ID:       1,
				IsShared: true,
			},
		},
	}).Error)
	CheckError(db.Create(&entity.Device{
		DeviceID:       deviceID,
		UserID:         1,
		LastLoggedInAt: time.Now().UTC(),
	}).Error)
	CheckError(db.Create(&entity.File{
		FileName:      "example.txt",
		RelativePath:  "/sample",
		IsDirectory:   true,
		LatestVersion: 1,
		Checksum:      "checksumallalal",
		NamespaceID:   1,
		UpdatedAt:     time.Now().UTC(),
		FileVersions: []entity.FileVersion{
			{
				ID:            1,
				DeviceID:      deviceID,
				VersionNumber: 1,
				UpdatedAt:     time.Now().UTC(),
				Device: entity.Device{
					DeviceID:       deviceID,
					UserID:         1,
					LastLoggedInAt: time.Now().UTC(),
				},
				Blocks: []entity.Block{
					{
						BlockID:       1,
						FileVersionID: 1,
						BlockOrder:    0,
					},
					{
						BlockID:       2,
						FileVersionID: 1,
						BlockOrder:    1,
					},
				},
			},
		},
	}).Error)
}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
