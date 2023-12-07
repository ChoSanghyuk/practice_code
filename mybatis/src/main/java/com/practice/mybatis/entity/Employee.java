package com.practice.mybatis.entity;

import lombok.Data;

@Data
public class Employee {
    private Long id;
    private String name;
    private int deptNo;

    public Employee(){

    }
    public Employee(Long id, String name, int deptNo){
        this.id = id;
        this.name = name;
        this.deptNo = deptNo;
    }
}
