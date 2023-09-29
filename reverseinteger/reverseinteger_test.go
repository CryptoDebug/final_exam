package reverseinteger

import (
    "testing"
)

func TestReverseInteger(t *testing.T) {
    tests := []struct {
        input    int
        expected int
    }{
        {123, 321},
        {-456, -654},
        {987, 789},
        {505, 505},
        {-999, -999},
        {12345, 54321},
        {-6789, -9876},
        {98765, 56789},
        {-123456, -654321},
        {12, 21},
        {-34, -43},
        {56, 65},
        {-78, -87},
        {-11, -11},
        {22, 22},
        {-33, -33},
        {44, 44},
        {-55, -55},
    }

    for testnumber, test := range tests {
        result := ReverseInteger(test.input)
        if result != test.expected{
            t.Logf("Test %d failed :( Your answer: %v ||| The good answer: %d", testnumber, result, test.expected)
        } else {
            t.Logf("Test %d passed!", testnumber)
        }
    }
}