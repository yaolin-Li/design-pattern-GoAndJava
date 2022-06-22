package com.basePriciples.Liskov_Substitution_Principle;

public class Square implements Quadrilateral{

    private double side;

    public double getSide() {
        return side;
    }

    public void setSide(double side) {
        this.side = side;
    }

    @Override
    public double getLength() {
        // TODO Auto-generated method stub
        return side;
    }

    @Override
    public double getWidth() {
        // TODO Auto-generated method stub
        return side;
    }

}
