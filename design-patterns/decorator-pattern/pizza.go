package main

type IPizza interface {
	getPrice() int
}

type VegPizza struct{}

func (vp *VegPizza) getPrice() int {
	return 199
}

type TomatoTopping struct {
	pizza IPizza
}

func (tt *TomatoTopping) getPrice() int {
	currPrice := tt.pizza.getPrice()
	return currPrice + 10
}

type CheeseTopping struct {
	pizza IPizza
}

func (ct *CheeseTopping) getPrice() int {
	currPrice := ct.pizza.getPrice()
	return currPrice + 20
}
