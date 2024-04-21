package entity

import "time"

type Namespace struct {
	ID        int64 `gorm:"primaryKey"`
	OwnerID   int64 `gorm:"index"`
	IsShared  bool
	CreatedAt time.Time
	UpdatedAt time.Time

	Files []File `gorm:"foreignKey:NamespaceID"`
}
