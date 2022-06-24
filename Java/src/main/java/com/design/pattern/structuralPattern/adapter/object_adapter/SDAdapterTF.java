package com.design.pattern.structuralPattern.adapter.object_adapter;
//通过继承和实现，来做到两者的互相调用
public class SDAdapterTF implements SDCard {

    private TFCard tfCard;

    public SDAdapterTF(TFCard tfCard) {
        this.tfCard = tfCard;
    }

    @Override
    public String readSD() {
        // TODO Auto-generated method stub
        System.out.println("adapter read tf card");
        return tfCard.readTF();
    }

    @Override
    public void writeSD(String msg) {
        // TODO Auto-generated method stub
        System.out.println("adapter write tf card");
        tfCard.writeTF(msg);
    }
}
