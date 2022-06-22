package com.builderPattern.FactoryPattern.CoffeDemo.before;

public abstract class Coffe {

    public abstract String getName();

    public void addsugar() {
        System.out.println("加糖");
    }

    public void addMike() {
        System.out.println("加奶");
    }
}
