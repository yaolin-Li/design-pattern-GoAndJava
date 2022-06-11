package IIGoversion

import "fmt"

type MyDoor struct{}

func (m MyDoor) AntiTheft() {
	fmt.Println("防盗")
}

func (m MyDoor) FireProof() {
	fmt.Println("防火")
}