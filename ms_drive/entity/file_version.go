package entity

import (
	"time"
)

type FileVersion struct {
	ID            int64  `gorm:"primaryKey"`
	FileID        int64  `gorm:"index;not null;"`
	DeviceID      string `gorm:"index;not null;"`
	VersionNumber int64
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Device Device  `gorm:"foreignKey:DeviceID;references:DeviceID"`
	Blocks []Block `gorm:"foreignKey:FileVersionID"`
}
