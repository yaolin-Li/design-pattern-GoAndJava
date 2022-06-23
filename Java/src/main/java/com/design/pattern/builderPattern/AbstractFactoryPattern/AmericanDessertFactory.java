package com.design.pattern.builderPattern.AbstractFactoryPattern;


public class AmericanDessertFactory implements DessertFactory{

    @Override
    public Dessert creatDessert() {
        // TODO Auto-generated method stub
        return new MatchaMousse();
    }

    @Override
    public Coffee createCoffee() {
        // TODO Auto-generated method stub
        return new AmericanCoffee();
    }

}
