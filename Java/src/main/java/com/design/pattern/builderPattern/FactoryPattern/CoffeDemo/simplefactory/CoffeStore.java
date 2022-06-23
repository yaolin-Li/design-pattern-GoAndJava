package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.simplefactory;


public class CoffeStore {
    public Coffe orderCoffee(String type) {
        SimpleCoffeeFactory factory = new SimpleCoffeeFactory();
        Coffe c = factory.createCoffee(type);
        c.addMike();
        c.addsugar();
        return c;
    }
}
