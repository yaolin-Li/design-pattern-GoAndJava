package classadapter

import "fmt"

type SDAdapterTF struct{
	tfcard TFCardImpl
}
func (s SDAdapterTF)readSD() string {
	fmt.Println("适配器阅读TF卡")
	return s.tfcard.readTF()
}

func (s SDAdapterTF)writeSD(msg string){
	fmt.Println("适配器写入TF卡",msg)
	s.tfcard.writeTF(msg)
}