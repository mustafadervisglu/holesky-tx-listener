package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	ethClient *ethclient.Client
}

func NewClient(url string) (*Client, error) {
	client, err := ethclient.DialContext(context.Background(), url)

	if err != nil {
		return nil, err
	}
	return &Client{ethClient: client}, nil
}
