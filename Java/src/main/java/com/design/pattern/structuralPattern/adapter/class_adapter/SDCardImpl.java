package com.design.pattern.structuralPattern.adapter.class_adapter;

public class SDCardImpl implements SDCard{

    @Override
    public String readSD() {
        // TODO Auto-generated method stub
        return "SDCard read msg: hello";
    }

    @Override
    public void writeSD(String msg) {
        // TODO Auto-generated method stub
        System.out.println("SDCard write msg : "+msg);
    }
    
}
