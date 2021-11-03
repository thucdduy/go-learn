package main

import "fmt"

type hp struct {
}

func (hp *hp) printFile() {
	fmt.Println("printing by an HP printer.")
}
