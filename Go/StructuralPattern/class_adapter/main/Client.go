package main

import (
	. "design_pattern/StructuralPattern/class_adapter"
	"fmt"
)

func main() {
	computer := new(Computer)
	msg := computer.ReadSD(new(SDCardImpl))
	fmt.Println(msg)
	fmt.Println("+++++++++++++++++++++")
	msg1 := computer.ReadSD(new(SDAdapterTF))
	fmt.Println(msg1)
}