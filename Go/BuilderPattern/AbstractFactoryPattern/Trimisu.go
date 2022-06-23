package AbstractFactoryPattern

import "fmt"

type Trimisu struct {

}

func (t *Trimisu)Show() {
	fmt.Println("提拉米苏")
}