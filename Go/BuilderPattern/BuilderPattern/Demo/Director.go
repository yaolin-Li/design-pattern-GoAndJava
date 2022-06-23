package BuilderPattern


type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director{
	return &Director{
		builder: builder,
	}
}

func (d *Director)Construct() Bike{
	d.builder.BuilderFrame()
	d.builder.BuilderSeat()
	return d.builder.CreateBike()
}