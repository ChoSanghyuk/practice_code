package com.practice.FileValue.bean;

import com.practice.FileValue.obj.PlainObject;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertTrue;

public class PlainObjectTest {

    PlainObject PlainObject = new PlainObject();

    @Test
    public void getYmlStringValueTest(){ assertTrue(PlainObject.getName().equals("Yaml File")); }

    @Test
    public void getYmlIntValueTest(){ assertTrue(PlainObject.getCount() == 10); }
}
