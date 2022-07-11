package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{0, 2, 3, 4, 5}
	printSlice(s)

	// slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop it first two values.
	s = s[2:]
	printSlice(s)
}
