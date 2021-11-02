/*

- component interface 	--> iPizza
- concrete component 1 	--> ChickenPizza
- concrete component 2  --> TomatoPizza

- concrete decorator 1 --> PepperDecorator
	+ component properties
- concrete decorator 2 --> CheeseDecorator
	+ component properties

, trong đó:
 + concreate decorator x là tính năng thêm vào cho component đang có.

*/

package main

import (
	"fmt"
	d "test/decorator"
)

func main() {
	tomato := &d.TomatoPizza{}
	chicken := &d.ChickenPizza{}

	fmt.Println(tomato.DoPizza())
	fmt.Println(chicken.DoPizza())

	pepperDecorator := d.NewPepperDecorator(chicken)
	fmt.Println(pepperDecorator.DoPizza())
	pepperDecorator.SetPizza(tomato)
	fmt.Println(pepperDecorator.DoPizza())

	cheeseDecorator := d.NewCheeseDecorator(chicken)
	fmt.Println(cheeseDecorator.DoPizza())
	cheeseDecorator.SetPizza(tomato)
	fmt.Println(cheeseDecorator.DoPizza())
}
