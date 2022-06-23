package main

import (
	"design_pattern/BuilderPattern/FactoryPattern/simpleFactory/demo1"
	"fmt"
)

func main() {
	
	americanCoffee, _ := demo1.OrderCoffee("american")
	lattern, _ := demo1.OrderCoffee("lattern")
	printDetails(americanCoffee)
    printDetails(lattern)
}

func printDetails(g demo1.AbstractCoffee) {
    fmt.Printf("coffee: %s \n", g.GetName())
    fmt.Println()
}