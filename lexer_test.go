package bfgo

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestLexer(t *testing.T) {
	source, err := os.Open(filepath.Join("testdata", "helloworld.bf"))
	if err != nil {
		t.Fatalf("failed to open source file: %v", err)
	}
	defer source.Close()

	lexer := NewLexer(bufio.NewReader(source))

	ops := make([]Op, 0)

	for {
		op, err := lexer.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("lexer error: %v", err)
		}
		ops = append(ops, op)
	}

	if len(ops) != 68 {
		t.Errorf("expected 68 operations, got %d", len(ops))
	}
}
