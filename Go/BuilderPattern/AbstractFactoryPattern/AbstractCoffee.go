package AbstractFactoryPattern

type abstractCoffee interface {
	GetName() string
	addsuger()
	addMike()
	toString() string
}