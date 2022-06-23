package BuilderPattern

type Builder interface {
	BuilderFrame()
	BuilderSeat()
	CreateBike() Bike
}