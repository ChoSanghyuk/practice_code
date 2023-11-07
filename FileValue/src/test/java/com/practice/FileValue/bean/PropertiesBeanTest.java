package com.practice.FileValue.bean;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
public class PropertiesBeanTest {

    @Autowired
    PropertiesBean propertiesBean;

    @Test
    public void getPropertiesStringValueTest(){
        assertTrue(propertiesBean.getName().equals("Properties File"));
    }
    @Test
    public void getPropertiesIntValueTest(){
        assertTrue(propertiesBean.getCount() == 5);
    }
}
