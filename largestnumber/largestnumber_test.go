package largestnumber

import (
	"testing"
)

func TestLargestNumber(t *testing.T) {
    tests := []struct {
        input    []int
        expected int
    }{
        {[]int{1, 2, 3}, 3},
        {[]int{-1, -5, -3}, -1},
        {[]int{}, 0},
        {[]int{5001, 200, 3500}, 5001},
        {[]int{150, 800, 4254}, 4254},
        {[]int{999999, 999991, 9990}, 999999},
        {[]int{123456789, 987654321, 555555555}, 987654321},
        {[]int{876543210, 111111111, 999999999}, 999999999},
        {[]int{123456789, 987654321, 444444444}, 987654321},
    }

    for testnumber, test := range tests {
        result := LargestNumber(test.input)
        if result != test.expected {
            t.Logf("Test %d failed :( Your answer: %v | The good answer: %d", testnumber, result, test.expected)
        } else {
            t.Logf("Test %d passed!", testnumber)
        }
    }
}