package FactoryMethod

type CoffeeStore struct {
	factory CoffeeFactory
}

func (c *CoffeeStore)SetFactory(factory CoffeeFactory) {
	c.factory = factory
}

func (c *CoffeeStore)OrderCoffee() abstractCoffee{
	coffee := c.factory.createCoffee()
	coffee.addMike()
	coffee.addsuger()
	return coffee
}