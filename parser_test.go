package bfgo

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	source, err := os.Open(filepath.Join("testdata", "helloworld.bf"))
	if err != nil {
		t.Fatalf("failed to open source file: %v", err)
	}
	defer source.Close()

	ops, err := Parse(bufio.NewReader(source))
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	if len(ops) != 71 {
		t.Errorf("expected 71 operations, got %d", len(ops))
	}
}

func TestParserMatchedBrackets(t *testing.T) {
	input := "[+]"
	ops, err := Parse(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	if len(ops) != 3 {
		t.Fatalf("expected 3 operations, got %d", len(ops))
	}

	// ops[0] should be JumpFwd pointing to ops[2]
	if ops[0].kind != OpJumpFwd {
		t.Errorf("expected OpJumpFwd at position 0, got %s", ops[0].String())
	}
	if ops[0].operand != 2 {
		t.Errorf("expected JumpFwd operand to be 2, got %d", ops[0].operand)
	}

	// ops[2] should be JumpBack pointing to ops[0]
	if ops[2].kind != OpJumpBack {
		t.Errorf("expected OpJumpBack at position 2, got %s", ops[2].String())
	}
	if ops[2].operand != 0 {
		t.Errorf("expected JumpBack operand to be 0, got %d", ops[2].operand)
	}
}

func TestParserNestedBrackets(t *testing.T) {
	input := "[[+]]"
	ops, err := Parse(bufio.NewReader(strings.NewReader(input)))
	if err != nil {
		t.Fatalf("parser error: %v", err)
	}

	if len(ops) != 5 {
		t.Fatalf("expected 5 operations, got %d", len(ops))
	}

	// ops[0] should be JumpFwd pointing to ops[4]
	if ops[0].operand != 4 {
		t.Errorf("expected outer JumpFwd operand to be 4, got %d", ops[0].operand)
	}

	// ops[1] should be JumpFwd pointing to ops[3]
	if ops[1].operand != 3 {
		t.Errorf("expected inner JumpFwd operand to be 3, got %d", ops[1].operand)
	}

	if ops[3].operand != 1 {
		t.Errorf("expected inner JumpBack operand to be 1, got %d", ops[3].operand)
	}

	if ops[4].operand != 0 {
		t.Errorf("expected outer JumpBack operand to be 0, got %d", ops[4].operand)
	}
}

func TestParserUnmatchedOpenBracket(t *testing.T) {
	input := "[++"
	_, err := Parse(bufio.NewReader(strings.NewReader(input)))
	if err == nil {
		t.Fatalf("expected error for unmatched '[', got nil")
	}
}

func TestParserUnmatchedCloseBracket(t *testing.T) {
	input := "++]"
	_, err := Parse(bufio.NewReader(strings.NewReader(input)))
	if err == nil {
		t.Fatalf("expected error for unmatched ']', got nil")
	}
}
