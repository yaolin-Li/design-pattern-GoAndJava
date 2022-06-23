package BuilderPattern

type Bike struct {
	frame string
	seat string
}

func (b Bike)GetFrame() string{
	return b.frame
}

func (b *Bike)SetFrame(frame string) {
	b.frame = frame
}

func (b Bike)GetSeat() string {
	return b.seat
}

func (b *Bike)SetSeat(seat string) {
	b.seat = seat
}