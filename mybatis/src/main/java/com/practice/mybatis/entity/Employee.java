package com.practice.mybatis.entity;

public class Employee {
    private int id;
    private String name;
    private int deptNo;

    public Employee(){

    }
    public Employee(int id, String name, int deptNo){
        this.id = id;
        this.name = name;
        this.deptNo = deptNo;
    }
}
