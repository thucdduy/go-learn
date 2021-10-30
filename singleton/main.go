package main

import (
	"fmt"
	"testsingleton/singleton"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)

	for i := 0; i < 10; i++ {
		fmt.Println(s1.AddOne())
		fmt.Println(s1.AddOne())
	}
}
