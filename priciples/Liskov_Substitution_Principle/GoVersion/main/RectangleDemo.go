package main

import (
	LSgoVersion "LSGoVersion"
)

func main() {
	r := new(LSgoVersion.Rectangle)
	r.SetLength(20)
	r.SetWidth(10)
	resize(r)
	printLengthAndWidth(r)

	s := new(LSgoVersion.Square)
	s.SetSide(10)
	//resize(s) 无法使用，假如正方形还继承着长方形的话，这就会卡死，现在不会了
	printLengthAndWidth(s)
}

func resize(r *LSgoVersion.Rectangle) {
	for r.GetWidth() <= r.GetLength() {
		r.SetWidth(r.GetWidth() + 1)
	}
}

func printLengthAndWidth(q LSgoVersion.Quadrilateral) {
	println(int(q.GetLength()))
	println(int(q.GetWidth()))
}