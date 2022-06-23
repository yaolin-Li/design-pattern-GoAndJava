package AbstractFactoryPattern

type DessertFactory interface{
	CreateCoffee() abstractCoffee
	CreateDessert() abstractDessert
}