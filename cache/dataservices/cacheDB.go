package data_service

import (
	"errors"
	"sync"

	"github.com/mastik5h/LLD/cache/models"
)

type InMemoryDB[K models.Key, V models.Value] struct {
	db         map[K]V
	dbMetadata map[K]models.ValueMetadata
	lock       []sync.Mutex
}

func returnSamePointer[K models.Key, V models.Value](obj models.IDb[K, V]) *models.IDb[K, V] {
	// Convert the object to a pointer to MyInterface
	ptr := &obj
	return ptr
}

func InitializeInMemoryDB[K models.Key, V models.Value](size int) *models.IDb[K, V] {
	val := &InMemoryDB[K, V]{
		db:         make(map[K]V, size),
		dbMetadata: make(map[K]models.ValueMetadata, size),
		lock:       make([]sync.Mutex, 2),
	}
	return returnSamePointer[K, V](val)
}

func (db *InMemoryDB[K, V]) Get(key K) *V {
	if v, ok := db.db[key]; ok {
		return &v
	}
	return nil
}

func (db *InMemoryDB[K, V]) GetMetadata(key K) *models.ValueMetadata {
	if v, ok := db.dbMetadata[key]; ok {
		return &v
	}
	return nil
}

func (db *InMemoryDB[K, V]) Set(key K, value V) {
	db.lock[0].Lock()
	db.db[key] = value
	db.lock[0].Unlock()
}

func (db *InMemoryDB[K, V]) SetMetadata(key K, md models.ValueMetadata) {
	db.lock[1].Lock()
	db.dbMetadata[key] = md
	db.lock[1].Unlock()
}

func (db *InMemoryDB[K, V]) Size() int {
	return len(db.db)
}

func (db *InMemoryDB[K, V]) Remove(key K) error {
	if _, ok := db.db[key]; ok {
		delete(db.db, key)
		delete(db.dbMetadata, key)
	} else {
		return errors.New("unknown key asked to delete")
	}
	return nil
}
