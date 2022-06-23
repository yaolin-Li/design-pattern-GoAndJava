package Dependency_Inversion

import "fmt"

type MyDisk struct{}

func (m MyDisk) Save(str string) {
	fmt.Println("使用硬盘存储数据为：", str)
}

func (m MyDisk) Get() string {
	fmt.Println("使用硬盘读取数据")
	return "数据"
}
