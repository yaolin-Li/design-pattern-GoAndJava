package Liskov_Substitution_Principle

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) GetLength() float64 {
	return r.length
}

func (r Rectangle) GetWidth() float64 {
	return r.width
}

func (r *Rectangle) SetLength(length float64) {
	r.length = length
}

func (r *Rectangle) SetWidth(width float64) {
	r.width = width
}
