package com.practice.mybatis.dao;

import com.practice.mybatis.entity.Employee;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.mapping.Environment;
import org.apache.ibatis.session.Configuration;
import org.apache.ibatis.session.SqlSession;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
@Slf4j
public class MyDaoTest {

    @Autowired
    MyDao dao;

    @Test
    public void retrieveEmployeeByIdTest(){
        Employee employee = dao.retrieveEmployeeById(84455L);
        log.info("{}", employee);
        assertTrue(employee.getName().equals("조상혁"));
    }

}
