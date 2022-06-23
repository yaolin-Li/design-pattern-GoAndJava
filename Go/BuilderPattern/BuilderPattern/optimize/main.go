package main

import "fmt"

func main() {
	phone := new(Builder).
			SetCpu("intel").
			SetMainboard("华硕").
			SetMemory("金士顿").
			SetScreen("三星").
			build()
	fmt.Println(phone.toString())
}