package com.example.aop.entity;

import java.util.Objects;

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

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getDeptNo() {
        return deptNo;
    }

    public void setDeptNo(int deptNo) {
        this.deptNo = deptNo;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Employee employee = (Employee) o;
        return id == employee.id && deptNo == employee.deptNo && Objects.equals(name, employee.name);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, name, deptNo);
    }
}
