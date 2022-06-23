package FactoryMethod

type abstractCoffee interface {
	GetName() string
	addsuger()
	addMike()
	toString() string
}