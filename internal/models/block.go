package model

type Block struct {
	Number     uint64 `gorm:"primaryKey"`
	Hash       string
	ParentHash string
	TimeStamp  uint64
}
