package main

import (
	. "design_pattern/BuilderPattern/BuilderPattern/Demo"
	"fmt"
)

func main() {
	director := NewDirector(new(OfoBuilder))
	bike := director.Construct()
	fmt.Println(bike.GetFrame())
	fmt.Println(bike.GetSeat())
}
