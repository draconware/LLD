package data_service

import (
	"fmt"

	"github.com/mastik5h/LLD/cache/algorithms"
	"github.com/mastik5h/LLD/cache/models"
)

type LRU_EvicitionPoliy[K models.Key] struct {
	LRU_DB algorithms.LinkedList[K]
}

func (ep *LRU_EvicitionPoliy[K]) Init() {
	ep.LRU_DB = algorithms.InitializeLinkedList[K]()
}

func (ep *LRU_EvicitionPoliy[K]) Evict() K {
	key := ep.LRU_DB.GetFirstNode()
	fmt.Println("Going to evict this key: ", key)
	ep.LRU_DB.RemoveNode(key)
	return key
}

func (ep *LRU_EvicitionPoliy[K]) Update(key K) {
	ep.LRU_DB.RemoveNode(key)
	ep.LRU_DB.AddNode(key)
}
