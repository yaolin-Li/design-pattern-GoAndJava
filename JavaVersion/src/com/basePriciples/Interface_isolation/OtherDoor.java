package com.basePriciples.Interface_isolation;

public class OtherDoor implements AntiTheft, FireProof{
    public void antiTheft(){
        System.out.println("防盗");
    }

    public void fireProof(){
        System.out.println("防火");
    }
}
