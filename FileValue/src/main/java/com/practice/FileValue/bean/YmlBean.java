package com.practice.FileValue.bean;

import com.practice.FileValue.Factory.YmlLoadFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.PropertySource;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Component;

@Component
@PropertySource(value = "/application1.yml", factory = YmlLoadFactory.class)
public class YmlBean {

    private String name;
    @Value("${YmlBean.field.count}")
    private int count;

    @Autowired
    public YmlBean(Environment env){
        this.name = env.getProperty("YmlBean.field.name", String.class);
    }

    public String getName() { return name;}
    public int getCount(){
        return count;
    }

}
