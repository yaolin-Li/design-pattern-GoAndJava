package com.builderPattern.SingletonPattern.EagerSingleton.Demo3;
public class Client{
    public static void main(String[] args) {
        EagerSingleton instance1 = EagerSingleton.INSTANCE;
        EagerSingleton instance2 = EagerSingleton.INSTANCE;
        System.out.println(instance1 == instance2);
    }
}
