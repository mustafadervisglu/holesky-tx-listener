package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type Listener struct {
	db     *gorm.DB
	client *ethclient.Client
}

func NewListener(client *ethclient.Client, db *gorm.DB) *Listener {
	return Listener{
		db:     db,
		client: client,
	}
}
