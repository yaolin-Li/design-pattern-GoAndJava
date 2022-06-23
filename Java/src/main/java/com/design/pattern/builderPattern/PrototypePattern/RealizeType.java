package com.design.pattern.builderPattern.PrototypePattern;
//浅克隆
public class RealizeType implements Cloneable{

    public RealizeType() {
        System.out.println("创建完成");
    }

    @Override
    protected Object clone() throws CloneNotSupportedException {
        System.out.println("复制成功");
        return super.clone();
    }
}
