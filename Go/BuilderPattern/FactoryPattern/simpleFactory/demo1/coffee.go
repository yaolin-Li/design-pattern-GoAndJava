package demo1

import "fmt"

type coffee struct {
	name string
}

func (c *coffee) addsuger() {
	fmt.Println("加糖")
}

func (c *coffee) addMike() {
	fmt.Println("加奶")
}

func (c *coffee) GetName() string {
	return c.name
}

func (c *coffee) toString() string {
	return "类型：" + c.GetName()
}
