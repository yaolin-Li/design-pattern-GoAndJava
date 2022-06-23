package Law_of_Demeter

type Start struct{
	name string
}

func NewStartByName(name string) *Start{
	return &Start{
		name: name,
	}
}

func (s Start) getName() string{
	return s.name
}