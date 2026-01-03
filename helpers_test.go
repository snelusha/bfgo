package bfgo

import "testing"

func TestIsBrainfuckChar(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'+', true},
		{'-', true},
		{'<', true},
		{'>', true},
		{'.', true},
		{',', true},
		{'[', true},
		{']', true},
		{'a', false},
		{'1', false},
		{'$', false},
	}

	for _, test := range tests {
		result := isBrainfuckChar(test.input)
		if result != test.expected {
			t.Errorf("isBrainfuckChar(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsRepeatable(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
	}{
		{'+', true},
		{'-', true},
		{'<', true},
		{'>', true},
		{'.', true},
		{',', false},
		{'[', false},
		{']', false},
		{'a', false},
		{'1', false},
		{'$', false},
	}

	for _, test := range tests {
		result := isRepeatable(test.input)
		if result != test.expected {
			t.Errorf("isRepeatable(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
