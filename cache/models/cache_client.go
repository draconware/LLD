package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"github.com/mastik5h/LLD/cache/helpers"
)

type Value interface {
	int | string | float64
}
type ValueMetadata struct {
	ExpiredAt time.Time
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
	Db             *IDb[K, V]
	WokerPool      WorkerPool
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
	keyChannel := make(chan K, 1)        // Buffered channel to avoid goroutine leak
	valueChannel := make(chan V, 1)      // Buffered channel to avoid goroutine leak
	errorChannel := make(chan string, 1) // Buffered channel for errors
	var wg sync.WaitGroup

	hashKey := cc.getHashKey(key)
	workerID := cc.WokerPool.GetWorkerIdForKey(hashKey)

	wg.Add(1)
	go func(workerID int) {
		defer wg.Done()
		cc.initiateWorkerForGet(workerID, keyChannel, valueChannel, errorChannel)
	}(workerID)

	keyChannel <- key
	close(keyChannel) // Close the key channel when you're done sending values

	wg.Wait()
	close(valueChannel)
	close(errorChannel)

	return <-valueChannel, errors.New(<-errorChannel) // Return the received value and error
}

func (cc *CacheClient[K, V]) SetEntry(key K, value V, time_to_live *int) error {
	value_md := ValueMetadata{
		ExpiredAt: helpers.GetTimeFromCurrentTime(time_to_live, cc.DefaultTTL),
	}
	fmt.Println("currentTime: ", time.Now().String(), ", expiredtime: ", value_md.ExpiredAt.String())
	key_channel := make(chan K, 1)
	value_channel := make(chan V, 1)
	value_md_channel := make(chan ValueMetadata, 1)
	error_channel := make(chan string, 1)
	var wg sync.WaitGroup

	hash_key := cc.getHashKey(key)
	workerId := cc.WokerPool.GetWorkerIdForKey(hash_key)
	wg.Add(1)
	go func(worker_id int) {
		defer wg.Done()
		cc.initiateWorkerForSet(worker_id, key_channel, value_channel, value_md_channel, error_channel)
	}(workerId)

	key_channel <- key
	value_channel <- value
	value_md_channel <- value_md

	close(key_channel)
	close(value_channel)
	close(value_md_channel)

	wg.Wait()

	close(error_channel)

	return errors.New(<-error_channel)
}

func (cc *CacheClient[K, V]) initiateWorkerForGet(wid int, key_channel <-chan K, value_channel chan<- V, errDetails_channel chan<- string) {
	fmt.Printf("Worker %d has picked up a task.\n", wid)
	var key K
	if k, ok := <-key_channel; ok {
		key = k
	}
	value, err := cc.getEntry(key)
	fmt.Printf("Worker %d finished processing key: %v\n", wid, key)

	if value != nil {
		value_channel <- *value
	}

	if err != nil {
		errDetails_channel <- err.Error()
	}
}
func (cc *CacheClient[K, V]) initiateWorkerForSet(wid int, key_channel <-chan K, value_channel <-chan V, value_md_channel <-chan ValueMetadata, errDetails chan<- string) {
	var key K
	var value V
	var value_md ValueMetadata
	if k, ok := <-key_channel; ok {
		key = k
	}
	if v, ok := <-value_channel; ok {
		value = v
	}
	if v_md, ok := <-value_md_channel; ok {
		value_md = v_md
	}
	fmt.Printf("Worker %d has started up task for %v: %v\n", wid, key, value)

	e := cc.setEntry(key, value, value_md)
	if e != nil {
		errDetails <- e.Error()
	}
	fmt.Printf("Worker %d finished processing %v: %v\n", wid, key, value)
}

func (cc *CacheClient[K, V]) getEntry(key K) (*V, error) {
	if cc.Db == nil {
		return nil, errors.New("no data found. Something unusual")
	}

	db := *(cc.Db)
	value := db.Get(key)
	if value == nil {
		return nil, errors.New("value doesn't exist in cache")
	}

	if value_md := db.GetMetadata(key); value_md != nil {
		expiredAt := value_md.ExpiredAt
		if expiredAt.Before(time.Now()) {
			if err := db.Remove(key); err != nil {
				return nil, errors.New("error removing expired key")
			}
			cc.EvictionPolicy.Update(key, true)
			fmt.Printf("Key is expired. key: %v\n", key)
			fmt.Println("currentTime: ", time.Now().String(), ", expiredtime: ", expiredAt.String())
			return cc.getEntry(key)
		}
	}
	ep := cc.EvictionPolicy
	ep.Update(key, false)
	fmt.Printf("Successfully retrived %v:%v\n", key, *value)
	return value, nil
}

func (cc *CacheClient[K, V]) setEntry(key K, value V, value_md ValueMetadata) error {
	if cc.Db == nil {
		return errors.New("no Db exists for cache")
	}

	db := *(cc.Db)
	ep := cc.EvictionPolicy
	if db.Size() == cc.CacheSize {
		evictedKey := ep.Evict()
		if err := db.Remove(evictedKey); err != nil {
			log.Printf("Error removing evicted key: %v\n", err)
			return err
		}
		log.Println("Evicted key: ", evictedKey)
	}

	db.Set(key, value)
	db.SetMetadata(key, value_md)
	ep.Update(key, false)

	fmt.Printf("Successfully added %v: %v\n", key, value)
	return nil
}
func (cc *CacheClient[K, V]) getHashKey(key K) int {
	keyValue := reflect.ValueOf(key)
	if keyValue.Kind() == reflect.Int {
		return int(keyValue.Int())
	} else if keyValue.Kind() == reflect.String {
		hash_key := helpers.GenerateConsistentHashKey(keyValue.String())
		return int(hash_key % uint64(cc.WokerPool.WorkerCount))
	} else {
		return 0
	}
}
