package com.practice.mybatis.dao;

import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.mapping.Environment;
import org.apache.ibatis.session.Configuration;
import org.apache.ibatis.session.SqlSession;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

@Repository
@Slf4j
public class MyDao {

    private SqlSession session;

    @Autowired
    public MyDao(SqlSession session){
        this.session = session;
    }

    public SqlSession getSqlSession(){
        return session;
    }
}
