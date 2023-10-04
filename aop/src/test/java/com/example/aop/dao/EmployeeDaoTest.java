package com.example.aop.dao;

import com.example.aop.entity.Employee;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.transaction.annotation.Transactional;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
@Transactional
public class EmployeeDaoTest {

    @Autowired
    EmployeeDao employeeDao;

    @Test
    public void insertTest(){
        Employee employee = new Employee(80000, "김민준", 2);
        int cnt = employeeDao.insert(employee);
        assertTrue(cnt == 1);

    }
}
