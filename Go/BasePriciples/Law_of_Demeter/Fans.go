package Law_of_Demeter

type Fans struct {
	name string
}

func NewFansByName(name string) *Fans{
	return &Fans{
		name: name,
	}
}

func (f Fans) getName() string {
	return f.name
}