package BuilderPattern

type MobileBuilder struct{
	bike Bike
}

func (m *MobileBuilder)BuilderFrame() {
	m.bike.SetFrame("碳纤维车架")
}

func (m *MobileBuilder)BuilderSeat(){
	m.bike.SetSeat("真皮坐垫")
}

func (m MobileBuilder)CreateBike() Bike {
	return m.bike
}
