package com.design.pattern.builderPattern.PrototypePattern;

public class Client {
    public static void main(String[] args) throws CloneNotSupportedException {
        RealizeType realizeType = new RealizeType();

        RealizeType clone = (RealizeType) realizeType.clone();

        System.out.println(realizeType == clone);
    }
}
