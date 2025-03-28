package main

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("insufficient funds found")
	}
	w.balance -= amount
	return nil
}
