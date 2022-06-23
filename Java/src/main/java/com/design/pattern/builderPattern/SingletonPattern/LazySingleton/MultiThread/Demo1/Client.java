package com.builderPattern.SingletonPattern.LazySingleton.MultiThread.Demo1;

public class Client{
    public static void main(String[] args) {
        for(int i = 0; i < 5; i++) {
            Obj obj = new Obj();
            obj.start();
        }
    }
}

class Obj extends Thread {

    @Override
    public void run() {
        Singleton instance = Singleton.getInstance();
        System.out.println(instance.hashCode());
    }
}
