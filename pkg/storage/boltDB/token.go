package boltDB

import (
	"errors"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/cvenkman/telegramBot/pkg/storage"
)

/* реализует интерфейс TokenStorageI */
type TokenStorage struct {
	db *bolt.DB
}

func NewTokenStorage(boltDB *bolt.DB) *TokenStorage {
	return &TokenStorage{db: boltDB}
}

func (ts *TokenStorage) Save(chatID int64, token string, bucket storage.Bucket) error {
	return ts.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put(intTiBytes(chatID), []byte(token))
	})
}

func (ts *TokenStorage) Get(chatID int64, bucket storage.Bucket) (string, error) {
	var token string

	err := ts.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		token = string(b.Get(intTiBytes(chatID)))
		return nil
	})

	if err != nil {
		return "", nil
	}
	if token == "" {
		return "", errors.New("token not found, maybe not valid chat id")
	}
	return token, nil
}

func intTiBytes(num int64) []byte {
	return []byte(strconv.FormatInt(num, 10))
}