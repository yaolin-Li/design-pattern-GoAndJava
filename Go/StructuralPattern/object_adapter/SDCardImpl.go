package objectadapter

import "fmt"

type SDCardImpl struct {

}

func (s SDCardImpl) readSD() string {
	return "从SD卡中读取数据"
}

func (s SDCardImpl) writeSD(msg string) {
	fmt.Println("从SD卡中读取数据",msg)
}