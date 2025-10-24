# mini-eth-node

```markdown
# 🪙 Mini Ethereum Node

A lightweight Ethereum "light node" built in Go.  
It connects to a public Ethereum RPC endpoint, periodically fetches the latest block, and stores block data (number → hash) in a local LevelDB database — simulating how real nodes maintain blockchain state.

---

## 🚀 Features

- Connects to the Ethereum mainnet using the official [`go-ethereum`](https://github.com/ethereum/go-ethereum) SDK
- Periodically syncs the latest block via **JSON-RPC**
- Logs block number, hash, and transaction count
- Persists block mappings (`blockNumber → hash`) using **LevelDB**
- Gracefully shuts down on `Ctrl + C`

---

## 🧱 Project Structure
```

mini-eth-node/
├── go.mod
├── main.go
└── node/
├── client.go # RPC connection to Ethereum node
├── storage.go # LevelDB wrapper for saving blocks
└── sync.go # Periodic sync logic

````

---

## ⚙️ Installation

```bash
git clone https://github.com/akenbay/mini-eth-node.git
cd mini-eth-node
go mod tidy
````

Dependencies:

- Go ≥ 1.22
- Libraries:

  - `github.com/ethereum/go-ethereum`
  - `github.com/syndtr/goleveldb/leveldb`

---

## 🧩 Usage

```bash
# Run the node
go run main.go
```

Example output:

```
Starting mini Ethereum node... syncing every 10s.
Latest Block: 21312489 | Hash: 0xabc... | Txns: 142
Latest Block: 21312490 | Hash: 0xdef... | Txns: 189
Latest Block: 21312491 | Hash: 0x219... | Txns: 176
```

Stop it anytime with `Ctrl + C` — it will shut down gracefully.

---

## 💾 Data Storage

Blocks are stored in a local LevelDB folder:

```
data/
  block_21312489 -> 0xabc...
  block_21312490 -> 0xdef...
```

To read from it manually:

```go
val, _ := db.GetBlock([]byte("block_21312489"), nil)
fmt.Println("Hash:", string(val))
```

---

## 🧑‍💻 Author

**Aibar Kenbay**
