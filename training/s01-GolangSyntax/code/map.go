package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{1000.0, 2000.0}
	fmt.Println(m["Bell Labs"])

	m1 := map[string]Vertex{
		"Vietnix": {1.1, 2.2},
		"Google":  {3.3, 4.4},
	}
	fmt.Println(m1)

	delete(m1, "Google")
	fmt.Println(m1)

	v, ok := m1["Vietnix"]
	fmt.Println("The Value:", v, ". Present?", ok)
}
