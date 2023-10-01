package models

import (
	"errors"
	"fmt"
	"time"
)

type Value interface {
	int | string | float64
}

type Key interface {
	comparable
}

type CacheClient[K Key, V Value] struct {
	Id             string
	EvictionPolicy IEvictionPolicy[K]
	CacheSize      int
	DefaultTTL     time.Duration
	CreatedAt      time.Time
	ModifiedAt     time.Time
	Trash          bool
	Db             IDb[K, V]
}

type CacheClientResponse struct {
	CacheClientId string
	ErrorDetails  string
}

type EvictionPolicyType string

const (
	LRU_EvicitionPolicyType EvictionPolicyType = "LRU_EvicitionPolicyType"
	LFU_EvicitionPolicyType EvictionPolicyType = "LFU_EvicitionPolicyType"
)

func (cc *CacheClient[K, V]) GetEntry(key K) (value V, err error) {
	if cc.Db == nil {
		err = errors.New("no Db exists for cache")
		return
	}

	if v := cc.Db.Get(key); v != nil {
		value = *v
		ep := cc.EvictionPolicy
		ep.Update(key)
	} else {
		err = errors.New("no Key entry exists in cache")
		return
	}
	fmt.Println("Successfully retrived ", key, ":", value)
	return
}

func (cc *CacheClient[K, V]) SetEntry(key K, value V) error {
	if cc.Db == nil {
		return errors.New("no Db exists for cache")
	}
	ep := cc.EvictionPolicy
	if cc.Db.Size() == cc.CacheSize {
		evictedKey := ep.Evict()
		err := cc.Db.Remove(evictedKey)
		if err != nil {
			return err
		}
	}
	cc.Db.Set(key, value)
	ep.Update(key)
	fmt.Println("Successfully added ", key, ":", value)
	return nil
}
