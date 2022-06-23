package Dependency_Inversion

type Cpu interface{
	Run()
}

type Disk interface {
	Save(str string)
	Get() string
}

type Memory interface {
	Save()
}