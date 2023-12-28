package com.practice.mybatis.config;

import com.mysql.cj.xdevapi.SessionFactory;
import com.zaxxer.hikari.HikariDataSource;
import lombok.extern.slf4j.Slf4j;
import org.apache.ibatis.session.SqlSessionFactory;
import org.mybatis.spring.SqlSessionFactoryBean;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.Resource;
import org.springframework.core.io.support.PathMatchingResourcePatternResolver;

import javax.sql.DataSource;
import java.io.File;
import java.io.IOException;
import java.io.InputStream;
import java.net.URI;
import java.net.URL;

@Configuration
@Slf4j
public class Config {

    @Bean
    @ConfigurationProperties(prefix= "spring.datasource.datasource-local")
    public DataSource dataSourceLocal() {
        return DataSourceBuilder.create().build();
    }
    @Bean
    @ConfigurationProperties(prefix= "spring.datasource.datasource-test")
    public DataSource dataSourceTest() {
        return DataSourceBuilder.create().build();
    }

    @Bean
    public SqlSessionFactory sqlSessionFactoryLocal(@Qualifier("dataSourceLocal") DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factory = new SqlSessionFactoryBean();
        factory.setDataSource(dataSource);
        factory.setConfigLocation(new PathMatchingResourcePatternResolver().getResource("/mybatis-config.xml"));
        Resource[] resources = new PathMatchingResourcePatternResolver()
                .getResources("classpath:sql/*.xml");
        factory.setMapperLocations(resources);
        return factory.getObject();
    }

    @Bean
    public SqlSessionFactory sqlSessionFactoryTest(@Qualifier("dataSourceTest") DataSource dataSource) throws Exception {
        SqlSessionFactoryBean factory = new SqlSessionFactoryBean();
        factory.setDataSource(dataSource);
        factory.setConfigLocation(new PathMatchingResourcePatternResolver().getResource("/mybatis-config.xml"));
        Resource[] resources = new PathMatchingResourcePatternResolver()
                .getResources("classpath:sql/*.xml");
        factory.setMapperLocations(resources);
        return factory.getObject();
    }


}
