package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	s := primes[1:3]
	fmt.Println(s)
}
