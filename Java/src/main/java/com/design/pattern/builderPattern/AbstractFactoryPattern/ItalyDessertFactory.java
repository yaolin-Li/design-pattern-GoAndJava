package com.design.pattern.builderPattern.AbstractFactoryPattern;


public class ItalyDessertFactory implements DessertFactory{

    @Override
    public Dessert creatDessert() {
        // TODO Auto-generated method stub
        return new Trimisu();
    }

    @Override
    public Coffee createCoffee() {
        // TODO Auto-generated method stub
        return new LatteCoffee();
    }

}
