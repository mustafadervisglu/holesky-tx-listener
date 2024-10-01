package ethereum

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	model "holesxy-tx-listener/internal/models"
	"log"
	"math/big"
	"strings"
	"sync"
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
	SaveBlock(blocks []model.Block) error
	ProcessBlocks(batchSize uint64)
}

//	func (l *Listener) Start() {
//		latestBlock, err := l.client.BlockNumber(context.Background())
//		if err != nil {
//			log.Println(err)
//		}
//		block, err := l.client.BlockByNumber(context.Background(), big.NewInt(int64(latestBlock)))
//		l.SaveBlock(block)
//		log.Printf("Latest block (Number: %d) saved successfully.", block.NumberU64())
//	}
func (l *Listener) SaveBlock(blocks []model.Block) error {
	if len(blocks) == 0 {
		return nil
	}
	if err := l.db.Create(&blocks).Error; err != nil {
		log.Printf("blocks could not be saved: %v", err)
		return err
	}
	log.Printf("Blocks saved successfully: %d : %d ", blocks[0].Number, blocks[len(blocks)-1].Number)
	return nil
}

func (l Listener) ProcessBlocks(batchSize uint64) {
	var blocks []model.Block
	var lastSavedBlock model.Block
	var startBlock uint64
	err := l.db.Order("number DESC").First(&lastSavedBlock).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			startBlock = 1
		} else {
			log.Printf("Could not get last block: %v", err)
			return
		}
	} else {
		startBlock = lastSavedBlock.Number + 1
	}

	endBlock := startBlock + batchSize - 1

	maxWorker := 20
	sem := make(chan struct{}, maxWorker)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {

		wg.Add(1)
		sem <- struct{}{}
		go func(blockNumber uint64) {
			defer wg.Done()
			defer func() { <-sem }()
			block, err := l.client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
			if err != nil {
				if isBlockNotFoundError(err) {
					log.Printf("block %d not found: %v", blockNumber, err)
					return
				} else {
					log.Printf("block %d could not get: %v", blockNumber, err)
				}
			}
			savedBlock := model.Block{
				Number:     block.NumberU64(),
				Hash:       block.Hash().Hex(),
				ParentHash: block.ParentHash().Hex(),
				TimeStamp:  block.Time(),
				EventSaved: false,
			}
			mu.Lock()
			blocks = append(blocks, savedBlock)
			mu.Unlock()
		}(blockNumber)
	}
	wg.Wait()

	err = l.SaveBlock(blocks)
	if err != nil {
		log.Printf("error occurred while saving blocks: %v", err)
	}
}

func isBlockNotFoundError(err error) bool {
	return errors.Is(err, ethereum.NotFound) || strings.Contains(err.Error(), "not found")
}
