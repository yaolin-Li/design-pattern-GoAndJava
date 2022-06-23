package com.design.pattern.builderPattern.Builder_Pattern;

public class OfoBuilder extends Builder{

    @Override
    public void buildFrame() {
        // TODO Auto-generated method stub
        bike.setFrame("铝合金车架");
    }

    @Override
    public void buildSeat() {
        // TODO Auto-generated method stub
        bike.setSeat("橡胶车座");
    }

    @Override
    public Bike createBike() {
        // TODO Auto-generated method stub
        return bike;
    }
    
}
