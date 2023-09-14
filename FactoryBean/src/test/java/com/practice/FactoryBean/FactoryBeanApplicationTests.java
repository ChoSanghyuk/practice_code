package com.practice.FactoryBean;

import com.practice.FactoryBean.Message.Message;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.jupiter.api.Assertions.assertTrue;

@SpringBootTest
class FactoryBeanApplicationTests {

	@Autowired
	Message message;

	@Test
	void messageBeanTest() {
		String rtn = message.getText();
		assertTrue(rtn.equals("factory bean sample test"));
	}

}
