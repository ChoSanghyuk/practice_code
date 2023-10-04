package com.example.aop.service;


import com.example.aop.entity.Employee;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.dao.TransientDataAccessResourceException;
import org.springframework.test.annotation.Rollback;
import org.springframework.transaction.annotation.Transactional;

import java.sql.SQLException;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
public class EmployeeServiceTest {

    @Autowired
    EmployeeService employeeService;

    @Test
    @Transactional(readOnly = true)
    public void test1(){
        Employee employee = new Employee(80000, "김민준", 2);
        assertThrows(TransientDataAccessResourceException.class,() ->employeeService.addOne(employee));
    }

    @Test
    @Transactional(readOnly = true)
    public void test2(){
        Employee employee1 = new Employee(80000, "김민준", 2);
        employeeService.addOneIndependent(employee1);
        Employee employee2 = employeeService.getOne(80000);
        assertTrue(employee1.equals(employee2));
    }
}
