package demo1

type latteCoffee struct {
	coffee
}

func newLatteCoffee() AbstractCoffee {
	return &latteCoffee{
		coffee: coffee{
			name: "拿铁咖啡",
		},
	}
}