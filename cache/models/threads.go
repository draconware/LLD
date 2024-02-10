package models

type WorkerPool struct {
	WorkerCount int
}

func (wp WorkerPool) GetWorkerIdForKey(key int) int {
	return key % wp.WorkerCount
}
