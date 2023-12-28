package com.practice.mybatis.dao;

import com.practice.mybatis.entity.Employee;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.session.SqlSession;
import org.apache.ibatis.session.SqlSessionFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Repository;

@Repository
@Slf4j
public class TestDao {
    private SqlSession session;

    @Autowired
    public TestDao(@Qualifier("sqlSessionFactoryTest") SqlSessionFactory sqlSessionFactory){
        this.session = sqlSessionFactory.openSession();
    }

    public Employee retrieveEmployeeById(Long id){
        return session.selectOne("O.Employee.retrieveById", id);
    }
}
