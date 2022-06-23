package FactoryMethod

type LatteCoffeeFactory struct{}

func (a *LatteCoffeeFactory)createCoffee() abstractCoffee{
	return newLatteCoffee()
}