package com.design.pattern.structuralPattern.adapter.object_adapter;

public class TFCardImpl implements TFCard{

    @Override
    public String readTF() {
        // TODO Auto-generated method stub
        return "TFCard read msg: hello";
    }

    @Override
    public void writeTF(String msg) {
        // TODO Auto-generated method stub
        System.out.println("TFCard write msg : "+msg);
    }
    
}
