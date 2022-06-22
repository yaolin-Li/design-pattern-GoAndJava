package com.basePriciples.Law_of_Demeter;

public class Client {
    public static void main(String[] args) {
        Agent agent = new Agent();
        Start start = new Start("周杰伦");
        agent.setStart(start);
        Fans fans = new Fans("小明");
        agent.setFans(fans);
        Company company = new Company("希望传媒");
        agent.setCompany(company);

        agent.meeting();
        agent.business();
    }
}
