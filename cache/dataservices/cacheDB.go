package data_service

import (
	"errors"

	"github.com/mastik5h/LLD/cache/models"
)

type InMemoryDB[K models.Key, V models.Value] struct {
	db map[K]V
}

func InitializeInMemoryDB[K models.Key, V models.Value]() models.IDb[K, V] {
	return &InMemoryDB[K, V]{
		db: make(map[K]V),
	}
}

func (db *InMemoryDB[K, V]) Get(key K) *V {
	if v, ok := db.db[key]; ok {
		return &v
	}
	return nil
}

func (db *InMemoryDB[K, V]) Set(key K, value V) {
	db.db[key] = value
}

func (db *InMemoryDB[K, V]) Size() int {
	return len(db.db)
}

func (db *InMemoryDB[K, V]) Remove(key K) error {
	if _, ok := db.db[key]; ok {
		delete(db.db, key)
	} else {
		return errors.New("unknown key asked to delete.")
	}
	return nil
}
