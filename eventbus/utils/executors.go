package utils

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

type ExecutorPool struct {
	Workers          map[WorkerType]([]*Worker)
	ExecutorPoolSize int
	wg               *sync.WaitGroup
}

func NewExecutorPool(executorPoolSize int) *ExecutorPool {
	if executorPoolSize <= 0 {
		return nil
	}
	executorPool := &ExecutorPool{
		ExecutorPoolSize: executorPoolSize,
		wg:               &sync.WaitGroup{},
	}
	executorPool.Workers = initializeWorkers(executorPoolSize, executorPool.wg)
	return executorPool
}

func (ep *ExecutorPool) Submit(workerType WorkerType, task ITask) (string, error) {
	if ep == nil {
		return "", errors.New("executor pool is empty. Cannot submit Task to an empty pool of Workers")
	}
	if task == nil {
		return "", errors.New("task found empty. Cannot assign worker to an empty Task")
	}
	taskId := task.GetId()
	taskIdHash := calculateHash(taskId)
	workersCount := len(ep.Workers[workerType])
	workerId := taskIdHash % workersCount
	fmt.Println("workerType:  ", workerType, " workerID: ", workerId, " workersCount: ", workersCount)
	fmt.Println("Executor Pool picked ")
	worker := ep.Workers[workerType][workerId]
	ep.wg.Add(1)
	go worker.Execute(task)
	return strconv.Itoa(workerId), nil
}

type Worker struct {
	Type      WorkerType
	TaskQueue chan ITask
	wg        *sync.WaitGroup
}

type WorkerType int

const (
	HandlePublishEventWorker = iota
	HandlePushEventWorker
	HandlePullEventByTimestampWorker
	HandlePullEventByIdWorker
)

func initializeWorkers(noOfWorkers int, wg *sync.WaitGroup) map[WorkerType]([]*Worker) {
	workers := make(map[WorkerType]([]*Worker), 0)
	for i := 0; i < noOfWorkers/4; i++ {
		workers[HandlePushEventWorker] = append(workers[HandlePushEventWorker], newWorker(HandlePushEventWorker, wg))
	}
	for i := noOfWorkers / 4; i < noOfWorkers/2; i++ {
		workers[HandlePullEventByTimestampWorker] = append(workers[HandlePullEventByTimestampWorker], newWorker(HandlePullEventByTimestampWorker, wg))
	}
	for i := noOfWorkers / 2; i < (3*noOfWorkers)/4; i++ {
		workers[HandlePullEventByIdWorker] = append(workers[HandlePullEventByTimestampWorker], newWorker(HandlePullEventByIdWorker, wg))
	}
	for i := (3 * noOfWorkers) / 4; i < noOfWorkers; i++ {
		workers[HandlePublishEventWorker] = append(workers[HandlePublishEventWorker], newWorker(HandlePublishEventWorker, wg))
	}
	return workers
}

func newWorker(workerType WorkerType, wg *sync.WaitGroup) *Worker {
	worker := &Worker{
		Type:      workerType,
		TaskQueue: make(chan ITask, 1),
		wg:        wg,
	}
	return worker
}

func (w *Worker) Execute(task ITask) {
	defer w.wg.Done()
	task.Execute()
}

type ITask interface {
	GetId() string
	Execute()
	Wait() map[string]interface{}
}
