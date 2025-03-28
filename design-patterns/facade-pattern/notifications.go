package main

import "fmt"

type WalletNotifications struct{}

func newWalletNotifications() *WalletNotifications {
	return &WalletNotifications{}
}

func (wn *WalletNotifications) sendCreditBalanceNotification() {
	fmt.Println("Sending wallet credit notifications.")
}

func (wn *WalletNotifications) sendDebitBalanceNotification() {
	fmt.Println("Sending wallet debit notifications.")
}
