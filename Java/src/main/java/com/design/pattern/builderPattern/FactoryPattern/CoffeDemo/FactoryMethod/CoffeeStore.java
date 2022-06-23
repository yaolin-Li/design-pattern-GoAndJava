package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.FactoryMethod;


public class CoffeeStore {
    private CoffeeFactory factory;

    public void setFactory(CoffeeFactory factory) {
        this.factory = factory;
    }

    public Coffee orderCoffee() {
        Coffee c = factory.createCoffee();
        c.addMike();
        c.addSugar();
        return c;
    }
}
