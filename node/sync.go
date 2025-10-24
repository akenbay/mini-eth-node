package node

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func StartSync(client *ethclient.Client, db *DB, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			syncLatestBlock(client, db)
		}
	}
}

func syncLatestBlock(client *ethclient.Client, db *DB) {
	ctx := context.Background()
	block, err := client.BlockByNumber(ctx, nil)

	if err != nil {
		slog.Error("Error when fetching block:", "error", err)
		return
	}

	fmt.Printf("Latest Block: %v | Hash: %s | Txns: %d\n",
		block.Number(), block.Hash(), len(block.Transactions()))

	err = db.SaveBlock(block.NumberU64(), block.Hash().Hex())

	if err != nil {
		slog.Error("Error when saving block:", "error", err)
		return
	}
}
