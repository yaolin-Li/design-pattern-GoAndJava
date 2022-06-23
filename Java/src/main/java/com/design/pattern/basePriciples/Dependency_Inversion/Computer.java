package com.design.pattern.basePriciples.Dependency_Inversion;

public class Computer {
    private AbstractDisk disk;
    private AbstractCpu cpu;
    private AbstractMemory memory;

    public AbstractDisk getDisk() {
        return disk;
    }

    public void setDisk(AbstractDisk disk) {
        this.disk = disk;
    }

    public AbstractCpu getCpu() {
        return cpu;
    }

    public void setCpu(AbstractCpu cpu) {
        this.cpu = cpu;
    }

    public AbstractMemory getMemory() {
        return memory;
    }

    public void setMemory(AbstractMemory memory) {
        this.memory = memory;
    }

    public void run() {
        System.out.println("运行计算机");
        String data = disk.get();
        System.out.println("从硬盘上获取的数据是："+data);
        cpu.run();
        memory.save();
    }
}
