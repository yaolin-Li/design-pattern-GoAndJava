package main

import DIGoversion "design_pattern/BasePriciples/Dependency_Inversion"

func main() {
	cpu := new(DIGoversion.MyCpu)
	disk := new(DIGoversion.MyDisk)
	memory := new(DIGoversion.MyMemory)

	c := new(DIGoversion.Computer)
	c.SetCpu(cpu)
	c.SetDisk(disk)
	c.SetMemory(memory)
	c.Run()
}