package main

import (
	"fmt"
	"time"

	"github.com/mastik5h/LLD/cache/client"
	"github.com/mastik5h/LLD/cache/helpers"
)

func main() {
	fmt.Println("Welcome to cache service.")
	helpers.InitializeStringGenerator()

	cache, err := client.CreateCache("4", "LRU", "60")
	if err != "" {
		fmt.Println(err)
		return
	}
	err = client.SetEntry(cache, "Mayank", "Aggarwal", "")
	fmt.Println(err)
	err = client.SetEntry(cache, "Komal", "Goyal", "")
	fmt.Println(err)
	time.Sleep(3 * time.Second)

	value, verr := client.GetEntry(cache, "Komal")
	if verr != "" {
		fmt.Println("error", verr)
	} else {
		fmt.Println(value)
	}

	value, verr = client.GetEntry(cache, "Mukul")
	if verr != "" {
		fmt.Println("error", verr)
	} else {
		fmt.Println(value)
	}
}
