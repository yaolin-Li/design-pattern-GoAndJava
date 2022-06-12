package LODGoversion

type Company struct{
	name string
}

func NewCompanyByName(name string) *Company {
	return &Company{
		name: name,
	}
}

func (c Company) getName() string {
	return c.name
}