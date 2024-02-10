package services

import (
	"errors"
	"time"

	data_service "github.com/mastik5h/LLD/cache/dataservices"
	"github.com/mastik5h/LLD/cache/helpers"
	"github.com/mastik5h/LLD/cache/models"
)

func InitializeCacheClient[K comparable, V models.Value](cache_size int, evictionPolicy models.EvictionPolicyType, default_ttl int) (models.ICache[K, V], error) {
	cacheClient := &models.CacheClient[K, V]{
		Id:         helpers.GetUniqueIdString(),
		CacheSize:  cache_size,
		DefaultTTL: time.Duration(default_ttl) * time.Second,
		CreatedAt:  time.Now(),
		Trash:      false,
		Db:         data_service.InitializeInMemoryDB[K, V](cache_size),
		WokerPool: models.WorkerPool{
			WorkerCount: 11,
		},
	}
	if evictionPolicy == models.LRU_EvicitionPolicyType {
		cacheClient.EvictionPolicy = &data_service.LRU_EvicitionPoliy[K]{}
	} else if evictionPolicy == models.LFU_EvicitionPolicyType {
		return nil, errors.New("LFU based cache not implemented yet. Sorry for inconvinience")
	}
	cacheClient.EvictionPolicy.Init()
	return cacheClient, nil
}
