package decorator

type TomatoPizza struct{}

func (c *TomatoPizza) DoPizza() string {
	return "I am a tomato pizza"
}
