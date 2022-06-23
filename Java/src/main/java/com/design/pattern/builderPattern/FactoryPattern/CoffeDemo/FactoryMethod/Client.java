package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.FactoryMethod;

public class Client {
    public static void main(String[] args) {
        CoffeeStore cs = new CoffeeStore();
        cs.setFactory(new AmericanCoffeeFactory());
        Coffee coffee = cs.orderCoffee();
        System.out.println(coffee.getName());
    }
}
