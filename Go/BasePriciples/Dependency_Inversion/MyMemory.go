package Dependency_Inversion

import "fmt"

type MyMemory struct{}

func (m MyMemory) Save() {
	fmt.Println("使用内存条")
}