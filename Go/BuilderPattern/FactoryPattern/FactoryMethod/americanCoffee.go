package FactoryMethod

type americanCoffee struct {
	coffee
}

func newAmericanCoffee() abstractCoffee {
	return &americanCoffee{
		coffee: coffee{
			name: "美式咖啡",
		},
	}
}