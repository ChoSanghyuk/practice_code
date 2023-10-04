package com.example.aop.service;

import com.example.aop.dao.EmployeeDao;
import com.example.aop.entity.Employee;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Component
@Transactional(propagation = Propagation.REQUIRED)
public class EmployeeService {

    EmployeeDao employeeDao;

    @Autowired
    public void EmployeeService(EmployeeDao employeeDao){
        this.employeeDao = employeeDao;
    }

    public void addOne(Employee employee){
        employeeDao.insert(employee);
    }

    @Transactional(propagation = Propagation.REQUIRES_NEW)
    public void addOneIndependent(Employee employee){
        employeeDao.insert(employee);
    }

    public void addAll(List<Employee> employeeList){
        for(Employee employee : employeeList){
            employeeDao.insert(employee);
        }
    }

    public Employee getOne(int id){
        return employeeDao.select(id);
    }
}
