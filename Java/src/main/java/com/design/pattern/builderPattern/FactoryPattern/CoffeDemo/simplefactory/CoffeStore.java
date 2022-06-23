package com.builderPattern.FactoryPattern.CoffeDemo.simplefactory;

import com.builderPattern.FactoryPattern.CoffeDemo.before.Coffe;

public class CoffeStore {
    public Coffe orderCoffee(String type) {
        SimpleCoffeeFactory factory = new SimpleCoffeeFactory();
        Coffe c = factory.createCoffee(type);
        c.addMike();
        c.addsugar();
        return c;
    }
}
