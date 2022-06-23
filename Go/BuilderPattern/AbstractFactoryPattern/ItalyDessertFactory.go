package AbstractFactoryPattern

type ItalyDessertFactory struct{}

func (a *ItalyDessertFactory)CreateCoffee() abstractCoffee {
	return newLatteCoffee()
}

func (a *ItalyDessertFactory)CreatDessert() abstractDessert {
	return new(Trimisu)
}