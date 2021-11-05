package main

import "fmt"

func main() {

	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()
	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %q\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %q\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Type: %d\n", normalHouse.floor)

	fmt.Println()
	fmt.Printf("Igloo House Door Type: %q\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %q\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Type: %d\n", iglooHouse.floor)

}
