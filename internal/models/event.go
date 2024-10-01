package model

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID          uint   `gorm:"primaryKey"`
	BlockNumber uint64 `gorm:"index"`
	TxHash      string `gorm:"index"`
	Index       uint
	From        string `gorm:"index"`
	To          string `gorm:"index"`
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
