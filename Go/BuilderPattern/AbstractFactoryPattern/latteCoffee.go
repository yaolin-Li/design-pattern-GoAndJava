package AbstractFactoryPattern

type latteCoffee struct {
	coffee
}

func newLatteCoffee() abstractCoffee {
	return &latteCoffee{
		coffee: coffee{
			name: "拿铁咖啡",
		},
	}
}