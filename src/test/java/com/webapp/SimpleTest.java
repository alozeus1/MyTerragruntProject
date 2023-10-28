package com.webapp;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

public class SimpleTest {

    @Test
    public void simpleTest() {
        int expected = 5;
        int result = 3 + 2;
        assertEquals(expected, result, "3 + 2 should equal 5");
    }
}
