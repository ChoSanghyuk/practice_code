package com.practice.testConfig.config;

import com.practice.testConfig.sample.MyBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class AppConfig {

    @Bean
    public MyBean myBean(){
        return new MyBean("Prod Environment Bean");
    }
}
