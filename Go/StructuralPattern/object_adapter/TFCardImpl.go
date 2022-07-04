package objectadapter

import "fmt"

type TFCardImpl struct {

}

func (t TFCardImpl)readTF() string {
	return "从TF卡中读取信息"
}

func (t TFCardImpl)writeTF(msg string) {
	fmt.Println("从TF卡中写入信息",msg)
}