package main

import (
	. "design_pattern/BuilderPattern/AbstractFactoryPattern"
	"fmt"
)

func main() {
	factory := new(AmericanDessertFactory)
	coffee := factory.CreateCoffee()
	dessert := factory.CreatDessert()
	fmt.Println(coffee.GetName())
	dessert.Show()
}