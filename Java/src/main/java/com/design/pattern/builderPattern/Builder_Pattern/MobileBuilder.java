package com.design.pattern.builderPattern.Builder_Pattern;

public class MobileBuilder extends Builder{

    @Override
    public void buildFrame() {
        // TODO Auto-generated method stub
        bike.setFrame("碳纤维车架");
    }

    @Override
    public void buildSeat() {
        // TODO Auto-generated method stub
        bike.setSeat("真皮车座");
    }

    @Override
    public Bike createBike() {
        // TODO Auto-generated method stub
        return bike;
    }
    
}
