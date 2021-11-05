package main

type gun struct {
	name  string
	power int
}

func (g *gun) setName(n string) {
	g.name = n
}

func (g *gun) setPower(p int) {
	g.power = p
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power

}
