package com.builderPattern.FactoryPattern.CoffeDemo.simplefactory;

import com.builderPattern.FactoryPattern.CoffeDemo.before.AmericanCoffe;
import com.builderPattern.FactoryPattern.CoffeDemo.before.Coffe;
import com.builderPattern.FactoryPattern.CoffeDemo.before.LatteCoffe;

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
