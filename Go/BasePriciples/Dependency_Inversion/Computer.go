package Dependency_Inversion

import "fmt"

type Computer struct {
	disk   Disk
	cpu    Cpu
	memory Memory
}
//切记set的时候要用指针接收者
func (c *Computer) SetDisk(disk Disk) {
	c.disk = disk
}

func (c *Computer) SetCpu(cpu Cpu) {
	c.cpu = cpu
}

func (c *Computer) SetMemory(memory Memory) {
	c.memory = memory
}

func (c Computer) GetDisk() Disk {
	return c.disk
}

func (c Computer) GetCpu() Cpu {
	return c.cpu
}

func (c Computer) GetMemory() Memory {
	return c.memory
}

func (c Computer) Run() {
	fmt.Println("运行计算机")
	data := c.disk.Get()
	fmt.Println("从硬盘上获取的数据是：", data)
	c.cpu.Run()
	c.memory.Save()
}
