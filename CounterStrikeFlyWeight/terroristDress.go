package main

type terroristDress struct {
	color string
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

func (t *terroristDress) getColor() string {
	return t.color
}
