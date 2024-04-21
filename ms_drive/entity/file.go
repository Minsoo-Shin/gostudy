package entity

import (
	"time"
)

type File struct {
	ID            int64  `gorm:"primaryKey"`
	FileName      string `gorm:"uniqueIndex:file_uri_idx"`
	RelativePath  string `gorm:"uniqueIndex:file_uri_idx"`
	IsDirectory   bool
	LatestVersion int64
	Checksum      string
	NamespaceID   int64 `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Namespace    Namespace     `gorm:"foreignKey:NamespaceID"`
	FileVersions []FileVersion `gorm:"foreignKey:FileID"`
}
