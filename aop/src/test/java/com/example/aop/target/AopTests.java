package com.example.aop.target;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
public class AopTests {

    @Autowired
    Target target;

    @Test
    public void aopLoggingTest(){
        // Client가 method1를 호출할 때에는 부가기능이 실행됩니다.
        target.method1();
    }


    @Test
    public void aopNoLoggingTest(){
        // Client가 아닌 타깃의 내부 메소드에서 타 메소드를 호출한 경우에는 부가기능이 실행되지 않습니다.
        target.method2();
    }

}
