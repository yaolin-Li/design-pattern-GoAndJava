package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.FactoryMethod;


public class AmericanCoffeeFactory implements CoffeeFactory{
    public Coffee createCoffee() {
        return new AmericanCoffee();
    }
}
