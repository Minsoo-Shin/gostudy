package entity

type Block struct {
	BlockID       int64 `gorm:"primaryKey"`
	FileVersionID int64 `gorm:"index;not null"`
	BlockOrder    int
}
