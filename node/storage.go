package node

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type DB struct {
	db *leveldb.DB
}

func NewDatabase(path string) (*DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (d *DB) SaveBlock(num uint64, hash string) error {
	key := []byte(fmt.Sprintf("block_%d", num))
	return d.db.Put(key, []byte(hash), nil)
}

func (d *DB) GetBlock(num uint64) (string, error) {
	key := []byte(fmt.Sprintf("block_%d", num))
	data, err := d.db.Get(key, nil)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (d *DB) Close() {
	d.db.Close()
}
