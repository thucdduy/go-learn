package main

import "fmt"

type ledger struct {
}

// makeEntry()
func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txtType %s for amount %d", accountID, txnType, amount)
}
