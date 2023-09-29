package removeduplicates

import (
    "testing"
)
func areEqualSlices(slice1, slice2 []int) bool {
    if len(slice1) != len(slice2) {
        return false
    }

    for i := 0; i < len(slice1); i++ {
        if slice1[i] != slice2[i] {
            return false
        }
    }

    return true
}

func TestRemoveDuplicates(t *testing.T) {
    tests := []struct {
        input    []int
        expected []int
    }{
        {[]int{1, 2, 2, 3, 4, 4, 5}, []int{1, 2, 3, 4, 5}},
        {[]int{7, 7, 7, 7, 7}, []int{7}},
        {[]int{1, 3, 3, 5, 5, 7, 9}, []int{1, 3, 5, 7, 9}},
        {[]int{2, 4, 6, 8, 8, 10, 10}, []int{2, 4, 6, 8, 10}},
        {[]int{5, 5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
        {[]int{11, 13, 17, 19, 19, 19}, []int{11, 13, 17, 19}},
        {[]int{0}, []int{0}},
        {[]int{9, 8, 8, 7, 6, 9, 5, 4, 3, 2, 1, 0}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
        {[]int{10, 20, 30, 40, 50}, []int{10, 20, 30, 40, 50}},
    }

    for testnumber, test := range tests {
        result := RemoveDuplicates(test.input)
        if !areEqualSlices(result, test.expected) {
            t.Logf("Test %d failed :( Your answer: %v ||| The good answer: %d", testnumber, result, test.expected)
        } else {
            t.Logf("Test %d passed!", testnumber)
        }
    }
}