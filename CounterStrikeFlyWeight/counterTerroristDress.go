package main

type counterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "blue"}
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}
