package menu

import "math/rand"

type Egg struct {
	Quantity int
	Quality  int
}

func (e *Egg) Init(quantity int) {
	e.Quantity = quantity
	e.Quality = rand.Intn(100) //return random number in interval (0,100)
}

func (e *Egg) GetQuantity() int {
	return e.Quantity
}

func (e *Egg) GetQuality() int {
	return e.Quality
}

func (e *Egg) Crack() {
}

func (e *Egg) DiscardShell() {
}

func (e *Egg) Cook() {
}
