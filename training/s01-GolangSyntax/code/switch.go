package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("go run on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		// BSD
		// Windows, other...
		fmt.Printf("%s\n", os)
	}
}
