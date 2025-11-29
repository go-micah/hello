package main

import "testing"

// TestGetMessage tests the getMessage function.
// Test functions must start with "Test" and take *testing.T as a parameter.
func TestGetMessage(t *testing.T) {
	// Table-driven tests are a Go idiom for testing multiple cases.
	// We define a slice of test cases, each with input and expected output.
	tests := []struct {
		name     string // description of the test case
		input    string // input to getMessage
		expected string // expected output
	}{
		{
			name:     "empty name returns default greeting",
			input:    "",
			expected: "Hello, World",
		},
		{
			name:     "provided name returns personalized greeting",
			input:    "Gopher",
			expected: "Hello, Gopher",
		},
		{
			name:     "name with spaces works correctly",
			input:    "Go Developer",
			expected: "Hello, Go Developer",
		},
	}

	// Loop through each test case
	for _, tt := range tests {
		// t.Run creates a subtest for each case
		t.Run(tt.name, func(t *testing.T) {
			got := getMessage(tt.input)
			if got != tt.expected {
				// t.Errorf reports a test failure with formatted output
				t.Errorf("getMessage(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
