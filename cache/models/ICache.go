package models

type IDb[K Key, V Value] interface {
	Get(key K) *V
	GetMetadata(key K) *ValueMetadata
	Set(key K, value V)
	SetMetadata(key K, value_md ValueMetadata)
	Remove(key K) error
	Size() int
}

type ICache[K Key, V Value] interface {
	GetEntry(key K) (V, error)
	SetEntry(key K, value V, time_to_live_in_minutes *int) error
}

type IEvictionPolicy[K Key] interface {
	Evict() K
	Update(key K, remove bool)
	Init()
}
