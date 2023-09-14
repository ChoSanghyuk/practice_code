package com.practice.FactoryBean.Message;

public class Message {

    String text;

    // private 생성자로 외부에서 생성자로 오브젝트 생성 불가
    private Message(String text){
        this.text = text;
    }

    public String getText(){
        return text;
    }
    // static method로 신규 객체 생성
    public static Message newMessage(String text){
        return new Message(text);
    }
}
