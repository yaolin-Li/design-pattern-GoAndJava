package DLGoversion

import "fmt"

type MyCpu struct{

}

func (m MyCpu) Run() {
	fmt.Println("使用Intel处理器")
}