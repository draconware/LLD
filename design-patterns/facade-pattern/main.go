package main

import (
	"fmt"
	"log"
)

func main() {
	walletFacade := NewWalletFacade("mayank", "007")
	fmt.Println()

	err := walletFacade.addMoneyToWallet("mayank", "007", 110)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	err = walletFacade.deductMoneyFromWallet("mayank", "007", 100)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	err = walletFacade.deductMoneyFromWallet("mayank", "007", 100)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
