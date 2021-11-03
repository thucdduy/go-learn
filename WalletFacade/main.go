package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println()

	walletFacade := newWalletFacade("NguyenPhuongHang", 1234)

	fmt.Println()

	if err := walletFacade.addMoneyToWallet("NguyenPhuongHang", 1234, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()

	if err := walletFacade.deductMoneyFromWallet("NguyenPhuongHang", 1234, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
