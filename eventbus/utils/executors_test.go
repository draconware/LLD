package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type MockTask struct {
	Id       string
	Args     map[string]interface{}
	Response chan map[string]interface{}
}

func NewMockTask(id string, args map[string]interface{}, response chan map[string]interface{}) *MockTask {
	return &MockTask{
		Id:       id,
		Args:     args,
		Response: response,
	}
}

func (mt *MockTask) Execute() {
	results := make(map[string]interface{}, 0)
	if val, ok := mt.Args["key1"]; ok {
		results["value1"] = val
	} else if val, ok = mt.Args["key2"]; ok {
		results["value2"] = val
	}
	mt.Response <- results
}

func (mt *MockTask) GetId() string {
	return mt.Id
}

func (mt *MockTask) Wait() map[string]interface{} {
	return nil
}

func TestNewExecutorPoolWithNonZeroPoolSize(t *testing.T) {
	poolSize := 10
	threadPool := NewExecutorPool(poolSize)
	if threadPool == nil {
		t.Fatalf("TestNewExecutorPoolWithNonZeroPoolSize failed for poolSize: %d", poolSize)
	}
}

func TestNewExecutorPoolWithZeroPoolSize(t *testing.T) {
	poolSize := 0
	threadPool := NewExecutorPool(poolSize)
	if threadPool != nil {
		t.Fatalf("TestNewExecutorPoolWithZeroPoolSize got created for poolSize: %d", poolSize)
	}
}

func TestExecutorPoolSubmit(t *testing.T) {
	IntializeGenerator(0)
	executorPool := NewExecutorPool(10)

	resp := make(chan map[string]interface{}, 1)
	args := map[string]interface{}{
		"key1": "mayank",
		"key2": "palak",
	}
	task := NewMockTask("task1", args, resp)

	executorPool.Submit(HandlePushEventWorker, task)
	response := <-resp
	if len(response) == 0 {
		t.Fatalf("TestExecutorPoolSubmit failed to send results on response channel")
	}
	for k, v := range response {
		fmt.Printf("key:%s, value:%s", k, v)
	}
}

func TestExecutorPoolSubmitForMultipleTaskWithSameTaskName(t *testing.T) {
	IntializeGenerator(0)
	executorPool := NewExecutorPool(10)
	args1 := map[string]interface{}{
		"key1": "mayank",
	}
	resp1 := make(chan map[string]interface{}, 1)
	task1 := NewMockTask("task", args1, resp1)

	args2 := map[string]interface{}{
		"key2": "palak",
	}
	resp2 := make(chan map[string]interface{}, 1)
	task2 := NewMockTask("task", args2, resp2)

	workerId1, _ := executorPool.Submit(HandlePushEventWorker, task1)
	response1 := <-resp1
	workerId2, _ := executorPool.Submit(HandlePushEventWorker, task2)
	response2 := <-resp2
	if len(response1) == 0 {
		t.Fatalf("TestExecutorPoolSubmit failed to send results on response1 channel")
	}
	if len(response2) == 0 {
		t.Fatalf("TestExecutorPoolSubmit failed to send results on response2 channel")
	}
	expectedResponse1 := map[string]interface{}{
		"value1": args1["key1"],
	}
	expectedResponse2 := map[string]interface{}{
		"value2": args2["key2"],
	}
	if !reflect.DeepEqual(response1, expectedResponse1) {
		t.Fatalf("TestExecutorPoolSubmit failed for task1 as expected is not equal to actual. Expected: %v, Actual: %v", expectedResponse1, response1)
	}
	if !reflect.DeepEqual(expectedResponse2, response2) {
		t.Fatalf("TestExecutorPoolSubmit failed for task2 as expected is not equal to actual. Expected: %v, Actual: %v", expectedResponse2, response2)
	}
	if workerId1 != workerId2 {
		t.Fatalf("TestExecutorPoolSubmit failed to assign workers as expected. task1WorkerId: %s, task2WorkerId: %s", workerId1, workerId2)
	}
}

func TestExecutorPoolSubmitForMultipleTaskWithDifferentTaskName(t *testing.T) {
	IntializeGenerator(0)
	executorPool := NewExecutorPool(10)
	args1 := map[string]interface{}{
		"key1": "mayank",
	}
	resp1 := make(chan map[string]interface{}, 1)
	task1 := NewMockTask("task1", args1, resp1)

	args2 := map[string]interface{}{
		"key2": "palak",
	}
	resp2 := make(chan map[string]interface{}, 1)
	task2 := NewMockTask("task2", args2, resp2)

	workerId1, _ := executorPool.Submit(HandlePushEventWorker, task1)
	response1 := <-resp1
	workerId2, _ := executorPool.Submit(HandlePushEventWorker, task2)
	response2 := <-resp2
	if len(response1) == 0 {
		t.Fatalf("TestExecutorPoolSubmit failed to send results on response1 channel")
	}
	if len(response2) == 0 {
		t.Fatalf("TestExecutorPoolSubmit failed to send results on response2 channel")
	}
	expectedResponse1 := map[string]interface{}{
		"value1": args1["key1"],
	}
	expectedResponse2 := map[string]interface{}{
		"value2": args2["key2"],
	}
	if !reflect.DeepEqual(response1, expectedResponse1) {
		t.Fatalf("TestExecutorPoolSubmit failed for task1 as expected is not equal to actual. Expected: %v, Actual: %v", expectedResponse1, response1)
	}
	if !reflect.DeepEqual(expectedResponse2, response2) {
		t.Fatalf("TestExecutorPoolSubmit failed for task2 as expected is not equal to actual. Expected: %v, Actual: %v", expectedResponse2, response2)
	}
	if workerId1 == workerId2 {
		t.Fatalf("TestExecutorPoolSubmit failed to assign workers as expected. task1WorkerId: %s, task2WorkerId: %s", workerId1, workerId2)
	}
}
