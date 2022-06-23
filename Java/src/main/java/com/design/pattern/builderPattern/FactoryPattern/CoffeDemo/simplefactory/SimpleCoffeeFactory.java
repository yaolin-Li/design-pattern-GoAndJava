package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.simplefactory;

public class SimpleCoffeeFactory {

    public Coffe createCoffee(String type) {
        Coffe c = null;
        if ("american".equals(type)) {
            c = new AmericanCoffe();
        } else if ("latte".equals(type)) {
            c = new LatteCoffe();
        }
        return c;
    }
}
