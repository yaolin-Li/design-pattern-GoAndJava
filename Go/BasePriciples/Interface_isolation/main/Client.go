package main

import IIGoversion "design_pattern/BasePriciples/Interface_isolation"

func main() {
	door := new(IIGoversion.MyDoor)

	door.AntiTheft()
	door.FireProof()
}