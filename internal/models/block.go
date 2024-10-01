package model

import (
	"time"
)

type Block struct {
	Number     uint64 `gorm:"primaryKey"`
	Hash       string
	ParentHash string
	TimeStamp  uint64
	EventSaved bool `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
