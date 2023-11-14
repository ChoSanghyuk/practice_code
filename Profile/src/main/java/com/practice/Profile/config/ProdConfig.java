package com.practice.Profile.config;

import com.practice.Profile.entity.MyBean;
import com.practice.Profile.loadFactory.YmlLoadFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Profile;
import org.springframework.context.annotation.PropertySource;

@Configuration
@Profile("prod")
@PropertySource(value = "/application.yml", factory = YmlLoadFactory.class)
public class ProdConfig {

    @Value("${custom.word}")
    String word;
    @Value("${custom.extra}")
    String extra;

    @Bean
    public MyBean myBean(){
        System.out.println("prod profile에서 빈을 생성합니다.");

        MyBean myBean = new MyBean(word, extra);
        return myBean;
    }
}
