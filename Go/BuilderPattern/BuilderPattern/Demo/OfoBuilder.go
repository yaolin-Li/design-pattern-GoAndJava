package BuilderPattern

type OfoBuilder struct{
	bike Bike
}

func (m *OfoBuilder)BuilderFrame() {
	m.bike.SetFrame("铝合金车架")
}

func (m *OfoBuilder)BuilderSeat(){
	m.bike.SetSeat("橡胶坐垫")
}

func (m OfoBuilder)CreateBike() Bike {
	return m.bike
}
