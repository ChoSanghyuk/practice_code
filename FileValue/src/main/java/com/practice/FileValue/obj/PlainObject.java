package com.practice.FileValue.obj;

import org.springframework.beans.factory.annotation.Value;
import org.yaml.snakeyaml.Yaml;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.util.Map;

public class PlainObject {
    private String name;
    private int count;

    public PlainObject(){
        try {
            Map<String, Object> config = new Yaml().load(new FileInputStream("src/main/resources/application1.yml"));
            Map<String, Object> yamlBeanConfig = (Map<String, Object>) config.get("YmlBean");
            Map<String, Object> yamlBeanFieldConfig = (Map<String, Object>) yamlBeanConfig.get("field");
            this.name = (String) yamlBeanFieldConfig.get("name");
            this.count = (int) yamlBeanFieldConfig.get("count");
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        }
    }

    public String getName() { return name;}
    public int getCount(){
        return count;
    }
}
