package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}
	{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "this is a test.",
			expected: []string{"this", "is", "a", "test."},
		},
		{
			input:    "foo bar",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "   leading",
			expected: []string{"leading"},
		},
		{
			input:    "trailing   ",
			expected: []string{"trailing"},
		},
		{
			input:    "  multiple   spaces   between  words  ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "one\ttwo\nthree",
			expected: []string{"one", "two", "three"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("wanted: %d, actual: %d", len(expected), len(actual))
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("wanted: %s, actual: %s", expectedWord, word)
		}
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}