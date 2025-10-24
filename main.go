package main

import (
	"fmt"
	"log/slog"
	"mini-eth-node/node"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	url := os.Getenv("ETH_API")
	client, err := node.NewClient(url)
	if err != nil {
		slog.Error("Error creating new client:", "error", err)
		return
	}
	defer client.Close()

	db, err := node.NewDatabase("data")
	if err != nil {
		slog.Error("Error creating new database:", "error", err)
		return
	}
	defer db.Close()

	fmt.Println("Starting mini Ethereum node... syncing every 5s.")
	go node.StartSync(client, db, 5*time.Second)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	fmt.Println("\nShutting down gracefully...")
}
