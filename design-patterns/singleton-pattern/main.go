package main

import "fmt"

func main() {
	for i := 0; i < 30; i++ {
		fmt.Println("trigerring go routine: ", i)
		go getClassAInstance()
	}
}
