package classadapter

type Computer struct{}

func (c Computer)ReadSD(sdCard SDCard) string{
	if sdCard == nil {
		panic("sd卡为空")
	}
	return sdCard.readSD()
}