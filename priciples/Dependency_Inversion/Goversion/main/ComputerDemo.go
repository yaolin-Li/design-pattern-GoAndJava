package main

import "DLGoversion"

func main() {
	cpu := new(DLGoversion.MyCpu)
	disk := new(DLGoversion.MyDisk)
	memory := new(DLGoversion.MyMemory)

	c := new(DLGoversion.Computer)
	c.SetCpu(cpu)
	c.SetDisk(disk)
	c.SetMemory(memory)
	c.Run()
}