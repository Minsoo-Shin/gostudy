package entity

import (
	"time"
)

// 사용자의 기기별 정보 및 로그인 기록 정보
type Device struct {
	ID             int64  `gorm:"primaryKey"`
	DeviceID       string `gorm:"uniqueIndex;not null"`
	UserID         int64  `gorm:"index;not null"`
	LastLoggedInAt time.Time
}
