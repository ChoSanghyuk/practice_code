package com.example.aop.target;

import org.springframework.stereotype.Component;

@Component
public class Target {

    public void method1(){
        System.out.println("Target Object do method1");
    }

    public void method2(){
        System.out.println("Target Object do method1 by method2");
        method1();
    }

}
