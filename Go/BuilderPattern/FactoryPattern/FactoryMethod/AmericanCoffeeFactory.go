package FactoryMethod

type AmericanCoffeeFactory struct{}

func (a *AmericanCoffeeFactory)createCoffee() abstractCoffee{
	return newAmericanCoffee()
}