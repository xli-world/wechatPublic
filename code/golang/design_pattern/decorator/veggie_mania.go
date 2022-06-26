package decorator

// veggeMania 具体的披萨
type veggeMania struct {
}

func (p *veggeMania) getPrice() int {
	return 15
}
