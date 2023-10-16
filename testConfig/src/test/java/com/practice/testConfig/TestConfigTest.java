package com.practice.testConfig;

import com.practice.testConfig.config.TestConfig;
import com.practice.testConfig.sample.MyBean;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.Import;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
@Import(TestConfig.class)
public class TestConfigTest {

    @Autowired
    private MyBean myTestBean;

    @Test
    public void beanTest(){
        assertTrue(myTestBean.getWords().equals("Test Environment Bean"));
    }
}
