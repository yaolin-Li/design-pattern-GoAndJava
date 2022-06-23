package com.basePriciples.Law_of_Demeter;

public class Agent {
    private Start start;
    private Fans fans;
    private Company company;
    public void setStart(Start start) {
        this.start = start;
    }
    public void setFans(Fans fans) {
        this.fans = fans;
    }
    public void setCompany(Company company) {
        this.company = company;
    }

    public void meeting() {
        System.out.println(start.getName()+"和粉丝"+fans.getName()+"见面");
    }

    public void business() {
        System.out.println(start.getName()+"和"+company.getName()+"洽谈");
    }
}
