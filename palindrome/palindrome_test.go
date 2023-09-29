package palindrome

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"racecar", true},
        {"kayak", true},
        {"radar", true},
        {"laval", true},
        {"rotor", true},
        {"bob", true},
        {"anna", true},
    }

    for testnumber, test := range tests {
        result := IsPalindrome(test.input)
        if result != test.expected {
            t.Logf("Test %d failed :( Your answer: %v | The good answer: %v", testnumber, result, test.expected)
        } else {
            t.Logf("Test %d passed!", testnumber)
        }
    }
}