package main

import (
	"fmt"
	"sync"
)

type ClassA struct{}

var lock = &sync.Mutex{}
var classAInstance *ClassA

func getClassAInstance() *ClassA {
	if classAInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if classAInstance == nil {
			fmt.Println("Creating single instance now.")
			classAInstance = &ClassA{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return classAInstance
}
