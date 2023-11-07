package com.practice.Profile.entity;

public class MyBean {

    String word;

    public MyBean(String env, String setting){
        this.word = String.format("%s with %s", env, setting);
        System.out.println(word);
    }
}
