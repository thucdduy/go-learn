package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := []bool{true, false, true, false}
	fmt.Println(b)

	c := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(c)

}
