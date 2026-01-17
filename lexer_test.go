package bfgo

import (
	"bufio"
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
			if err == ErrEOF {
				break
			}
			t.Fatalf("lexer error: %v", err)
		}
		ops = append(ops, op)
	}

	if len(ops) != 71 {
		t.Errorf("expected 71 operations, got %d", len(ops))
	}
}
