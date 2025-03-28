package main

import "fmt"

type Account struct {
	id string
}

func newAccount(accountId string) *Account {
	return &Account{
		id: accountId,
	}
}

func (a *Account) checkAccount(accountId string) error {
	if a.id == accountId {
		return nil
	}
	return fmt.Errorf("account name is incorrect")
}
