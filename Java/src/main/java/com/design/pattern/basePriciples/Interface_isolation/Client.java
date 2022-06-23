package com.design.pattern.basePriciples.Interface_isolation;

public class Client {
    public static void main(String[] args) {
        /**
         * 假设此时新的对象只需要防火和防盗不需要防水，那么直接继承SafetyDoor就有问题
         *
        MySafetyDoor my = new MySafetyDoor();

        my.antiTheft();
        my.fireProof();
        my.waterProof();
        */

        OtherDoor my = new OtherDoor();

        my.antiTheft();
        my.fireProof();
        //my.waterProof();
    }
}
