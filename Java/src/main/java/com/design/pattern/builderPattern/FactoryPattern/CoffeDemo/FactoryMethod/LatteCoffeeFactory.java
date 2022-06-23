package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.FactoryMethod;


public class LatteCoffeeFactory implements CoffeeFactory{
    public Coffee createCoffee() {
        return new LatteCoffee();
    }
}
