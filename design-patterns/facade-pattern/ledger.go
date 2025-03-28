package main

import "fmt"

type WalletLedger struct {
	ledger []string
}

func newWalletLedger() *WalletLedger {
	return &WalletLedger{
		ledger: make([]string, 0),
	}
}

func (wl *WalletLedger) makeEntry(accountId string, transactionType string, amount int) {
	wl.ledger = append(wl.ledger, fmt.Sprintf("Entry for account: %s with txnType: %s and amount: %d", accountId, transactionType, amount))
}
