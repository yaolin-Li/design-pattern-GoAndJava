package goversion

type Square struct {
	side float64
}

func (s Square) GetLength() float64 {
	return s.side
}

func (s Square) GetWidth() float64 {
	return s.side
}

func (s *Square) SetSide(side float64) {
	s.side = side
}
