package countvowels

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 2},
		{"world", 1},
		{"aeiou", 5},
		{"aeiou", 5},
		{"", 0},
		{"anticonstitutionnellement", 10},
		{"nantesynovcampus", 6},
		{"cryptographie", 5},
		{"algorithmique", 6},
		{"programmation", 5},
		{"langage", 3},
		{"logiciel", 4},
		{"informatique", 6},
		{"ordinateur", 5},
		{"internet", 3},
	}

	for testnumber, test := range tests {
		result := CountVowels(test.input)
		if result != test.expected {
			t.Logf("Test %d failed Your answer: %v | The good answer: %d", testnumber, result, test.expected)
		} else {
			t.Logf("Test %d passed!", testnumber)
		}
	}
}
