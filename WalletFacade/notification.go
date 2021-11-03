package main

import "fmt"

type notification struct {
}

// creditNotification()
func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

// debitNotification
func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
