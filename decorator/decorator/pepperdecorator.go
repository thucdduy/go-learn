package decorator

type PepperDecorator struct {
	pizza IPizza
}

func NewPepperDecorator(pizza IPizza) *PepperDecorator {
	return &PepperDecorator{
		pizza: pizza,
	}
}

func (p *PepperDecorator) DoPizza() string {
	pizzaType := p.pizza.DoPizza()
	return pizzaType + p.addFlavour()
}

func (c *PepperDecorator) SetPizza(pizza IPizza) {
	c.pizza = pizza
}

func (p *PepperDecorator) addFlavour() string {
	return " Pepper"
}
