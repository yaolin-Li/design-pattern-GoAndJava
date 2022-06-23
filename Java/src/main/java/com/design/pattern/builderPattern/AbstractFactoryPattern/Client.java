package com.design.pattern.builderPattern.AbstractFactoryPattern;

public class Client {
    public static void main(String[] args) {
        //ItalyDessertFactory factory = new ItalyDessertFactory();
        AmericanDessertFactory factory = new AmericanDessertFactory();
        Coffee coffee = factory.createCoffee();
        Dessert dessert = factory.creatDessert();
        System.out.println(coffee.getName());
        dessert.show();
    }
}
