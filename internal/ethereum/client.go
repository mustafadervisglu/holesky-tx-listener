package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(context.Background(), url)

	if err != nil {
		return nil, err
	}

	return client, nil
}
