package demo1


import (
	"strings"
)

func OrderCoffee(coffeeType string) (AbstractCoffee, error) {
	var coffee AbstractCoffee
	if strings.Compare(coffeeType,"american") == 0{
		coffee = newAmericanCoffee()
	} else if strings.Compare(coffeeType,"lattern") == 0{
		coffee = newLatteCoffee()
	}
	return coffee, nil
}