package com.design.pattern.builderPattern.FactoryPattern.CoffeDemo.simplefactory;


public class Client {
    public static void main(String[] args) {
        CoffeStore cs = new CoffeStore();
        Coffe c = cs.orderCoffee("latte");
        System.out.println(c.getName());
    }
}
