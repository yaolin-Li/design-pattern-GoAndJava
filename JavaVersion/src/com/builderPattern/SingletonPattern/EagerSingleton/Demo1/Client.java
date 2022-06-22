package com.builderPattern.SingletonPattern.EagerSingleton.Demo1;
public class Client{
    public static void main(String[] args) {
        EagerSingleton instance1 = EagerSingleton.getInstance();
        EagerSingleton instance2 = EagerSingleton.getInstance();
        System.out.println(instance1 == instance2);
    }
}
