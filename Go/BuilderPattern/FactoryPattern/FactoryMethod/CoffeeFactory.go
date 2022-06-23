package FactoryMethod

type CoffeeFactory interface{
	createCoffee() abstractCoffee
}