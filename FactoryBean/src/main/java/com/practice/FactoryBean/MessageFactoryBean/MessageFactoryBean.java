package com.practice.FactoryBean.MessageFactoryBean;

import com.practice.FactoryBean.Message.Message;
import org.springframework.beans.factory.FactoryBean;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;

@Component
public class MessageFactoryBean  implements FactoryBean<Message> {

    @Value("factory bean sample test")
    String text;


    @Override
    public Message getObject() throws Exception {
        return Message.newMessage(this.text);
    }

    @Override
    public Class<? extends Message> getObjectType() {
        return Message.class;
    }

    @Override
    public boolean isSingleton() {
        return false;
    }
}
