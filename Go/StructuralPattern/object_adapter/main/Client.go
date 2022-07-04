package main

import (
	. "design_pattern/StructuralPattern/object_adapter"
	"fmt"
)

func main() {
	computer := new(Computer)
	msg := computer.ReadSD(new(SDCardImpl))
	fmt.Println(msg)
	fmt.Println("+++++++++++++++++++++")
	msg1 := computer.ReadSD(NewSDAdapterTF(new(TFCardImpl)))
	fmt.Println(msg1)
}