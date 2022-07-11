package main

import "fmt"

func main() {
	a := [4]string{"AAA", "BBB", "CCC", "DDD"}
	b := a[0:2]
	c := a[1:3]
	fmt.Println(a, b, c)

	b[0] = "XXX"
	fmt.Println(a, b, c)
}
