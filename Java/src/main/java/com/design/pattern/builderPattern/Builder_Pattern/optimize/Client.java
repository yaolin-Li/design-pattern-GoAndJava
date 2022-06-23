package com.design.pattern.builderPattern.Builder_Pattern.optimize;

public class Client {
    public static void main(String[] args) {
        Phone phone = new Phone.Builder()
                        .cpu("intel")
                        .screen("三星")
                        .memory("金士顿")
                        .mainboard("华硕")
                        .build();

        System.out.println(phone.toString());
    }
}
