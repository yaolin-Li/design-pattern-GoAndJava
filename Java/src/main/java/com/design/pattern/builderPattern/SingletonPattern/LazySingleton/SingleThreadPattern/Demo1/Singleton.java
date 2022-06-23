package com.design.pattern.builderPattern.SingletonPattern.LazySingleton.SingleThreadPattern.Demo1;
//懒汉
public class Singleton {
    private Singleton(){}

    private static Singleton instance;

    public static Singleton getInstance() {
        if (instance == null) {
            instance = new Singleton();
        }
        return instance;
    }
}
