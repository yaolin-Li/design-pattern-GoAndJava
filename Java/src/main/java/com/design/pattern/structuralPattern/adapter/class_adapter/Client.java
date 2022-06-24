package com.design.pattern.structuralPattern.adapter.class_adapter;

public class Client {
    public static void main(String[] args) {
        Computer computer = new Computer();
        String msg = computer.readSD(new SDCardImpl());
        System.out.print(msg);

        System.out.println("\n+++++++++++++++++++++++++++++++");
        String msg1 = computer.readSD(new SDAdapterTF());
        System.out.print(msg1);
    }
}
