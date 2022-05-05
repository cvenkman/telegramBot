package boltdb

import (
	"github.com/cvenkman/telegramBot/pkg/storage"
)

type TokenStore struct {

}

func (ts *TokenStore) Save(chatID int64, token string, bucket storage.Bucket) error {
	return nil
}

func (ts *TokenStore) Get(chatID int64, bucket storage.Bucket) (string, error) {
	return "", nil
}
