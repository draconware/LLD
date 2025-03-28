package main

import "fmt"

type WalletFacade struct {
	account       *Account
	securityCode  *SecurityCode
	wallet        *Wallet
	notifications *WalletNotifications
	ledger        *WalletLedger
}

func NewWalletFacade(accountId, securityCode string) *WalletFacade {
	return &WalletFacade{
		account:       newAccount(accountId),
		securityCode:  newSecurityCode(securityCode),
		wallet:        newWallet(),
		notifications: newWalletNotifications(),
		ledger:        newWalletLedger(),
	}
}

func (wf *WalletFacade) addMoneyToWallet(accountId, securityCode string, amount int) error {
	fmt.Println("Credit money to wallet.")
	err := wf.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkSecurityCode(securityCode)
	if err != nil {
		return err
	}

	wf.wallet.creditBalance(amount)
	wf.notifications.sendCreditBalanceNotification()
	wf.ledger.makeEntry(accountId, "credit", amount)
	fmt.Printf("Amount: %d is credited to account: %s\n", amount, accountId)
	return nil
}

func (wf *WalletFacade) deductMoneyFromWallet(accountId, securityCode string, amount int) error {
	fmt.Println("Debit money from wallet.")
	err := wf.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkSecurityCode(securityCode)
	if err != nil {
		return err
	}

	err = wf.wallet.debitBalance(amount)
	if err != nil {
		return err
	}

	wf.notifications.sendDebitBalanceNotification()
	wf.ledger.makeEntry(accountId, "debit", amount)
	fmt.Printf("Amount: %d is debited from account: %s\n", amount, accountId)
	return nil
}
