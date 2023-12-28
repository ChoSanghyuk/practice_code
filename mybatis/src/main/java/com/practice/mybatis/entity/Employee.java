package com.practice.mybatis.entity;

import lombok.Data;
import lombok.Getter;
import lombok.ToString;

@Data
@ToString
public class Employee {
    private Long id;
    private String name;
    private int deptNo;

    public Employee(){
    }
}
