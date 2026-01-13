package bfgo

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestInterpreterBasicOperations(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		input    string
		expected string
	}{
		{
			name:     "simple output",
			program:  strings.Repeat("+", 65) + ".", // ASCII 65 = 'A'
			expected: "A",
		},
		{
			name:     "increment and output",
			program:  "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.",
			expected: "Hello World!",
		},
		{
			name:     "simple loop",
			program:  "+++++[>++<-]>.",
			expected: "\n",
		},
		{
			name:     "echo input",
			program:  ",.",
			input:    "X",
			expected: "X",
		},
		{
			name:     "pointer movement and output",
			program:  "++++>+++++>>+++<<<.>.",
			expected: string([]byte{4, 5}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			input := strings.NewReader(tt.input)
			reader := bufio.NewReader(strings.NewReader(tt.program))

			interp, err := NewInterpreter(reader, &output, input)
			if err != nil {
				t.Fatalf("failed to create interpreter: %v", err)
			}

			err = interp.Execute()
			if err != nil {
				t.Fatalf("execution failed: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected output %q, got %q", tt.expected, output.String())
			}
		})
	}
}

func TestInterpreterErrors(t *testing.T) {
	tests := []struct {
		name        string
		program     string
		expectError string
	}{
		{
			name:        "pointer underflow",
			program:     "<",
			expectError: "pointer underflow",
		},
		{
			name:        "pointer overflow",
			program:     strings.Repeat(">", 30000),
			expectError: "pointer overflow",
		},
		{
			name:        "input EOF",
			program:     ",",
			expectError: "input error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			input := strings.NewReader("")
			reader := bufio.NewReader(strings.NewReader(tt.program))

			interp, err := NewInterpreter(reader, &output, input)
			if err != nil {
				t.Fatalf("failed to create interpreter: %v", err)
			}

			err = interp.Execute()
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tt.expectError)
			}

			if !strings.Contains(err.Error(), tt.expectError) {
				t.Errorf("expected error containing %q, got %q", tt.expectError, err.Error())
			}
		})
	}
}

func TestInterpreterNestedLoops(t *testing.T) {
	// Nested loops: set cell 0 to 4, cell 1 to 12 (4*3)
	program := "++++[>+++[>+<-]<-]>>."
	expected := string([]byte{12})

	var output bytes.Buffer
	input := strings.NewReader("")
	reader := bufio.NewReader(strings.NewReader(program))

	interp, err := NewInterpreter(reader, &output, input)
	if err != nil {
		t.Fatalf("failed to create interpreter: %v", err)
	}

	err = interp.Execute()
	if err != nil {
		t.Fatalf("execution failed: %v", err)
	}

	if output.String() != expected {
		t.Errorf("expected output %q, got %q", expected, output.String())
	}
}

func TestInterpreterMemoryOperations(t *testing.T) {
	// Test increment and decrement
	program := "+++.--.+++++."
	expected := string([]byte{3, 1, 6})

	var output bytes.Buffer
	input := strings.NewReader("")
	reader := bufio.NewReader(strings.NewReader(program))

	interp, err := NewInterpreter(reader, &output, input)
	if err != nil {
		t.Fatalf("failed to create interpreter: %v", err)
	}

	err = interp.Execute()
	if err != nil {
		t.Fatalf("execution failed: %v", err)
	}

	if output.String() != expected {
		t.Errorf("expected output %q, got %q", expected, output.String())
	}
}

func TestInterpreterByteWrapping(t *testing.T) {
	tests := []struct {
		name     string
		program  string
		expected string
	}{
		{
			name:     "increment 255 wraps to 0",
			program:  strings.Repeat("+", 255) + "+.",
			expected: string([]byte{0}),
		},
		{
			name:     "decrement 0 wraps to 255",
			program:  "-.",
			expected: string([]byte{255}),
		},
		{
			name:     "multiple wraps",
			program:  strings.Repeat("+", 256) + strings.Repeat("+", 256) + ".",
			expected: string([]byte{0}),
		},
		{
			name:     "decrement wraps and increment back",
			program:  "-.+.",
			expected: string([]byte{255, 0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			input := strings.NewReader("")
			reader := bufio.NewReader(strings.NewReader(tt.program))

			interp, err := NewInterpreter(reader, &output, input)
			if err != nil {
				t.Fatalf("failed to create interpreter: %v", err)
			}

			err = interp.Execute()
			if err != nil {
				t.Fatalf("execution failed: %v", err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected output %q, got %q", tt.expected, output.String())
			}
		})
	}
}
