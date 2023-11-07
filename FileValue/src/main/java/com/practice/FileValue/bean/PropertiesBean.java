package com.practice.FileValue.bean;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.PropertySource;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Component;

@Component
@PropertySource("/application2.properties")
public class PropertiesBean {

    @Value("${PropertiesBean.field.name}")
    private String name;
    private int count;

    @Autowired
    public PropertiesBean(Environment env){
        this.count = env.getProperty("PropertiesBean.field.count", Integer.class);
    }

    public String getName(){
        return this.name;
    }

    public int getCount(){
        return this.count;
    }

}
