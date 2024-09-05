package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	model "holesxy-tx-listener/internal/models"
	"log"
	"math/big"
)

type Listener struct {
	db     *gorm.DB
	client *ethclient.Client
}

func NewListener(client *ethclient.Client, db *gorm.DB) *Listener {
	return &Listener{
		db:     db,
		client: client,
	}
}

type IListerner interface {
	Start()
	SaveBlock(block *types.Block)
}

func (l *Listener) Start() {
	latestBlock, err := l.client.BlockNumber(context.Background())
	if err != nil {
		log.Println(err)
	}
	block, err := l.client.BlockByNumber(context.Background(), big.NewInt(int64(latestBlock)))
	l.SaveBlock(block)
	log.Printf("Latest block (Number: %d) saved successfully.", block.NumberU64())
}

func (l *Listener) SaveBlock(block *types.Block) {
	saveBlock := model.Block{
		Number:     block.NumberU64(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		TimeStamp:  block.Time(),
	}
	if err := l.db.Create(&saveBlock).Error; err != nil {
		log.Printf("Failed to save block: %v", err)
	}
}
