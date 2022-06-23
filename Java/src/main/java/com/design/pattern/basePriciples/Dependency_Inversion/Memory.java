package com.design.pattern.basePriciples.Dependency_Inversion;

public class Memory implements AbstractMemory{

    public void save() {
        System.out.println("使用内存条");
    }
}
