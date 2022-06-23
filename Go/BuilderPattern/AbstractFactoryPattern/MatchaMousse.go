package AbstractFactoryPattern

import "fmt"

type MatchaMousse struct {
}

func (m *MatchaMousse)Show() {
	fmt.Println("抹茶慕斯")
}