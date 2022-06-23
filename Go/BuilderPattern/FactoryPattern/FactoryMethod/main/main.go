package main

import (
	. "design_pattern/BuilderPattern/FactoryPattern/FactoryMethod"
	"fmt"
)
func main() {
	cs := new(CoffeeStore)
	cs.SetFactory(new(AmericanCoffeeFactory))
	coffee := cs.OrderCoffee()
	fmt.Println(coffee.GetName())
}