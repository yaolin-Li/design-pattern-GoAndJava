package com.builderPattern.SingletonPattern.LazySingleton.MultiThread.Demo2;
//双重加锁饿汉
public class Singleton {
    private Singleton(){}

    private static volatile Singleton instance;

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized(Singleton.class){
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
