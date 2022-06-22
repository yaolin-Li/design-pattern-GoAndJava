package com.basePriciples.Interface_isolation;

public class MySafetyDoor implements SafetyDoor{
    public void antiTheft(){
        System.out.println("更好的防盗");
    }

    public void fireProof(){
        System.out.println("更好的防火");
    }

    public void waterProof(){
        System.out.println("更好的防水");
    }
}
