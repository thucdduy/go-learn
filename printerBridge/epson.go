package main

import "fmt"

type epson struct {
}

func (e *epson) printFile() {
	fmt.Println("printing by an EPSON printer")
}
