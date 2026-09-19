package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
	}
	
	for _, c := range cases {
		actual := cleanInput(c.input)
		fmt.Println(actual, len(actual))
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("The string did not match with lenses !")
			break
		}
		
		expectedSplit := c.expected

		if !(slices.Equal(expectedSplit, actual)) {
			
			t.Errorf("The string did not matched !")
		} else {
			fmt.Printf("ok")
		}
	}
}
