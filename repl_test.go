package main

import (
	"fmt"
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
		{
			input:    "helloworld",
			expected: []string{"helloworld"},
		},
		{
			input:    "hello_world",
			expected: []string{"hello_world"},
		},
		{
			input:    "hello WORLD",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		input := cleanInput(c.input)
		if len(input) != len(c.expected) {
			if len(input) != len(c.expected) {
				t.Errorf("For input '%s', expected %s but got %s. Length mismatch: expected %d, got %d",
					c.input,
					fmt.Sprintf("%#v", c.expected),
					fmt.Sprintf("%#v", input),
					len(c.expected),
					len(input))
			}
		}
		for i := range input {
			inputWord := input[i]
			expectedWord := c.expected[i]
			if inputWord != expectedWord {
				t.Errorf("For input '%s', inputWord '%s' does not match expectedWord '%s'", c.input, inputWord, expectedWord)
			}
		}
	}
}
