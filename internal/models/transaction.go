package model

type Transaction struct {
	Hash      string `gorm:"primaryKey"`
	BlockHash string
	From      string
	To        string
	Value     string
}
