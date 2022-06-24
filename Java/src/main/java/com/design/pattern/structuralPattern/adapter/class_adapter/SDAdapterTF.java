package com.design.pattern.structuralPattern.adapter.class_adapter;
//通过继承和实现，来做到两者的互相调用
public class SDAdapterTF extends TFCardImpl implements SDCard {

    @Override
    public String readSD() {
        // TODO Auto-generated method stub
        System.out.println("adapter read tf card");
        return readTF();
    }

    @Override
    public void writeSD(String msg) {
        // TODO Auto-generated method stub
        System.out.println("adapter write tf card");
        writeTF(msg);
    }
}
