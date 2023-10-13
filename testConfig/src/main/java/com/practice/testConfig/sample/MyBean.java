package com.practice.testConfig.sample;

public class MyBean {

    private String words;

    public MyBean(String words){
        System.out.println(words);
        this.words = words;
    }

    public String getWords(){
        return words;
    }
}
