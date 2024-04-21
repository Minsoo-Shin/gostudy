package entity

import "time"

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time

	Namespaces []Namespace `gorm:"foreignKey:OwnerID"`
	Devices    []Device    `gorm:"foreignKey:UserID"`
}
