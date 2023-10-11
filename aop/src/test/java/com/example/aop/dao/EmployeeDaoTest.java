package com.example.aop.dao;

import com.example.aop.entity.Employee;
import org.aspectj.lang.annotation.After;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.dao.EmptyResultDataAccessException;
import org.springframework.transaction.annotation.Transactional;

import static org.junit.jupiter.api.Assertions.assertThrows;
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

    @AfterAll
    static void RollbackValidationTest(@Autowired EmployeeDao employeeDao) {
        assertThrows(EmptyResultDataAccessException.class, () -> employeeDao.select(80000));
    }

}
