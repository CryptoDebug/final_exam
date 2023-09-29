package validparentheses

import (
    "testing"
)

func TestIsValidParentheses(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"()", true},
        {"()[]{}", true},
        {"(]", false},
        {"([)]", false},
        {"{[]}", true},
        {"{{}}", true},
        {"[({})]", true},
        {"[()]", true},
        {"[", false},
        {"(){}", true},
        {"[[]]", true},
        {"{[()]}", true},
        {"{[()]", false},
        {"()", true},
        {"({[]})", true},
        {"()", true},
        {"[(){}]", true},
        {"()", true},
        {"{[()]}", true},
        {"{[(])}", false},
    }

    for testnumber, test := range tests {
        result := IsValidParentheses(test.input)
        if result != test.expected{
            t.Logf("Test %d failed :( ! Your answer: %v ||| The good answer: %v", testnumber, result, test.expected)
        } else {
            t.Logf("Test %d passed!", testnumber)
        }
    }
}