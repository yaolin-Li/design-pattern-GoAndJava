package com.builderPattern.SingletonPattern.EagerSingleton.Demo1;

/**
 * 饿汉式
 * 静态变量创建类的对象
 */
public class EagerSingleton {

    private EagerSingleton() {}

    private static EagerSingleton instance = new EagerSingleton();

    public static EagerSingleton getInstance() {
        return instance;
    }
}
