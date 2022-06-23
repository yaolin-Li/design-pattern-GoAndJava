package AbstractFactoryPattern

type AmericanDessertFactory struct{}

func (a *AmericanDessertFactory)CreateCoffee() abstractCoffee {
	return newAmericanCoffee()
}

func (a *AmericanDessertFactory)CreatDessert() abstractDessert {
	return new(MatchaMousse)
}