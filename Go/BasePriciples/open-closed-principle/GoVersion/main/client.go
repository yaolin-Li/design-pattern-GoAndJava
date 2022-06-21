package main

import (
	goversion "GoVersion"
)

func main()  {
	input := new(goversion.SougouInput)
	skin := new(goversion.DefaultSkin)
	//SetSKin的参数是AbstractSkin，添加新的皮肤只需要实现这个接口就能被set
	input.SetSKin(skin)
	input.Display()
}