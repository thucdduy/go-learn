package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Tuesday?")
	today := time.Now().Weekday()
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorow.")
	case today + 2:
		fmt.Println("In 2 days.")
	default:
		fmt.Println("Too far away.")
	}
}
