package models

type IDb[K Key, V Value] interface {
	Get(key K) *V
	Set(key K, value V)
	Remove(key K) error
	Size() int
}

type ICache[K Key, V Value] interface {
	GetEntry(key K) (V, error)
	SetEntry(key K, value V) error
}

type IEvictionPolicy[K Key] interface {
	Evict() K
	Update(key K)
	Init()
}
