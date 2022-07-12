package main

import (
	"fmt"
	"strings"
)

func main() {

	board := [][]string{
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
		{"_", "_", "_", "_"},
	}

	board[0][1] = "X"
	board[1][2] = "O"
	board[3][3] = "X"
	board[2][1] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
