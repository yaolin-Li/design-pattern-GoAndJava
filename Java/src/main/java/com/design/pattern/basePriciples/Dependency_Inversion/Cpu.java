package com.design.pattern.basePriciples.Dependency_Inversion;

public class Cpu implements AbstractCpu{

    public void run() {
        System.out.println("使用Intel处理器");
    }
}
