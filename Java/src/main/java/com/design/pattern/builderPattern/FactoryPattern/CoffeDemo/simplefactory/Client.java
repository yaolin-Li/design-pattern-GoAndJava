package com.builderPattern.FactoryPattern.CoffeDemo.simplefactory;

import com.builderPattern.FactoryPattern.CoffeDemo.before.Coffe;
import com.builderPattern.FactoryPattern.CoffeDemo.before.CoffeStore;

public class Client {
    public static void main(String[] args) {
        CoffeStore cs = new CoffeStore();
        Coffe c = cs.orderCoffee("latte");
        System.out.println(c.getName());
    }
}
