package com.practice.FactoryBean.MessageFactoryBean;

import com.practice.FactoryBean.Message.Message;
import org.springframework.beans.factory.FactoryBean;

public class MessageFactoryBean  implements FactoryBean<Message> {

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
