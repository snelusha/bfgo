package bfgo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type Lexer struct {
	reader *bufio.Reader
}

func NewLexer(reader *bufio.Reader) *Lexer {
	return &Lexer{reader: reader}
}

func (l *Lexer) Next() (Op, error) {
	c, err := l.read()
	if err != nil {
		return Op{}, err
	}

	if !isBrainfuckChar(c) {
		return l.Next()
	}

	if isRepeatable(c) {
		operand, err := l.countRepeats(c)
		if err != nil {
			return Op{}, fmt.Errorf("counting repeats: %w", err)
		}

		op, ok := NewOpWithOperand(c, operand)
		if !ok {
			return Op{}, fmt.Errorf("invalid repeatable operation: %c", c)
		}

		return op, nil
	}

	op, ok := NewOp(c)
	if !ok {
		return Op{}, fmt.Errorf("invalid operation: %c", c)
	}

	return op, nil
}

func (l *Lexer) countRepeats(target rune) (uint, error) {
	count := uint(1)

	for {
		next, err := l.peek()
		if err != nil {
			if err == ErrEOF {
				break
			}
			return 0, err
		}

		if next != target {
			break
		}

		if _, err = l.read(); err != nil {
			return 0, err
		}

		count = count + 1
	}

	return count, nil
}

func (l *Lexer) peek() (rune, error) {
	c, err := l.read()
	if err != nil {
		return 0, err
	}

	if err := l.unread(); err != nil {
		return 0, fmt.Errorf("unreading after peek: %w", err)
	}

	return c, nil
}

func (l *Lexer) read() (rune, error) {
	c, _, err := l.reader.ReadRune()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return 0, ErrEOF
		}
		return 0, err
	}
	return c, nil
}

func (l *Lexer) unread() error {
	return l.reader.UnreadRune()
}
