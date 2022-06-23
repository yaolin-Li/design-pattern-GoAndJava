package com.design.pattern.builderPattern.FactoryPattern.config_factory;

public abstract class Coffee {

    public abstract String getName();

    public void addSugar() {
        System.out.println("加糖");
    }

    public void addMike() {
        System.out.println("加奶");
    }
}
