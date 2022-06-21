package goversion

type SougouInput struct {

}

var skin AbstractSKin

func (s SougouInput) SetSKin(a AbstractSKin) {
	skin = a
}

func (s SougouInput) Display() {
	skin.display()
}