package com.basePriciples.Dependency_Inversion;

public class ComputerDemo {
    public static void main(String[] args) {
        /**
         * 这样的写法无法修改disk/cpu/memory的类型，修改的话就必须修改computer类的代码
         * 此时应当将各种需要组合使用的东西进行抽象。
         * 这样一来，当有新的组件时，只需要继承抽象结构，并传递进computer类就好了，不需要修改computer类
        Disk disk = new Disk();
        Cpu cpu = new Cpu();
        Memory memory = new Memory();

        Computer c = new Computer();
        c.setCpu(cpu);
        c.setDisk(disk);
        c.setMemory(memory);

        c.run();
         */
        AbstractDisk disk = new Disk();
        AbstractCpu cpu = new Cpu();
        AbstractMemory memory = new Memory();

        Computer c = new Computer();
        c.setCpu(cpu);
        c.setDisk(disk);
        c.setMemory(memory);

        c.run();
    }
}
