package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.simplefactory;

public abstract class Coffe {

    public abstract String getName();

    public void addsugar() {
        System.out.println("加糖");
    }

    public void addMike() {
        System.out.println("加奶");
    }
}
