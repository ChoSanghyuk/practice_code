package com.practice.testConfig.config;

import com.practice.testConfig.sample.MyBean;
import org.springframework.boot.test.context.TestConfiguration;
import org.springframework.context.annotation.Bean;

@TestConfiguration
public class TestConfig {

    @Bean
    public MyBean myTestBean(){
        return new MyBean("Test Environment Bean");
    }
}
