package menu

type Chicken struct {
	Quantity int
}

func (e *Chicken) Init(quantity int) {
	e.Quantity = quantity
}

func (e *Chicken) GetQuantity() int {
	return e.Quantity
}

func (e *Chicken) CutUp() {
}

func (e *Chicken) Cook() {
}
