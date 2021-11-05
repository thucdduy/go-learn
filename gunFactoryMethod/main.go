package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetails(ak47)
	printDetails(maverick)

}

func printDetails(g iGun) {
	fmt.Printf("Gun: %q - Power: %d\n", g.getName(), g.getPower())
}
