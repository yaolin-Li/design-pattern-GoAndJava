package objectadapter

import "fmt"

type SDAdapterTF struct{
	tfCard TFCard
}

func NewSDAdapterTF(tfCard TFCard) *SDAdapterTF {
	return &SDAdapterTF{
		tfCard: tfCard,
	}
}
func (s SDAdapterTF)readSD() string {
	fmt.Println("适配器阅读TF卡")
	return s.tfCard.readTF()
}

func (s SDAdapterTF)writeSD(msg string){
	fmt.Println("适配器写入TF卡",msg)
	s.tfCard.writeTF(msg)
}