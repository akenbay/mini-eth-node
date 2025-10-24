package node

import "github.com/ethereum/go-ethereum/ethclient"

func NewClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)

	if err != nil {
		return nil, err
	}
	return client, nil
}
