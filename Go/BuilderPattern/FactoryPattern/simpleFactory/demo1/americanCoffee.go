package demo1

type americanCoffee struct {
	coffee
}

func newAmericanCoffee() AbstractCoffee {
	return &americanCoffee{
		coffee: coffee{
			name: "美式咖啡",
		},
	}
}