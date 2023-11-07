package com.practice.FileValue.bean;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
public class YmlBeanTest {

    @Autowired
    YmlBean ymlBean;

    @Test
    public void getYmlStringValueTest(){assertTrue(ymlBean.getName().equals("Yaml File")); }

    @Test
    public void getYmlIntValueTest(){  assertTrue(ymlBean.getCount() == 10); }
}
