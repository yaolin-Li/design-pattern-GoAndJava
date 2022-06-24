package com.design.pattern.structuralPattern.proxyMode.static_proxy;

public class ProxyPoint implements SellTickets{

    private TrainStation trainStation = new TrainStation();

    public void sell() {
        System.out.println("收取代理服务费用");
        trainStation.sell();
    }
}
