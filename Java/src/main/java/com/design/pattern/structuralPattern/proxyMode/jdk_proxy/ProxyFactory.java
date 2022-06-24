package com.design.pattern.structuralPattern.proxyMode.jdk_proxy;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

public class ProxyFactory {
    
    private TrainStation trainStation = new TrainStation();

    public SellTickets getProxyObject() {
        //类加载器+字节码对象+代理对象的调用处理程序
        SellTickets proxyObject = (SellTickets)Proxy.newProxyInstance(
            trainStation.getClass().getClassLoader(), 
            trainStation.getClass().getInterfaces(),
            new InvocationHandler() {
               @Override
               public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
                   // proxy:代理对象，和proxyObject是同一个，在invoke中基本不用
                   //method:对接口中的方法进行封装的method对象
                   //args:调用方法传入的参数
                   //返回值:就是方法的返回值
                    System.out.println("代收点收取费用(jdk)");
                    Object obj = method.invoke(trainStation, args);
                   return obj;
               } 
            });
        return proxyObject;
    }
}
