package bfgo

import (
	"bufio"
	"fmt"
	"io"
)

type Lexer struct {
	reader *bufio.Reader
}

func NewLexer(reader *bufio.Reader) *Lexer {
	return &Lexer{reader: reader}
}

func (this Lexer) Next() (Op, error) {
	c, err := this.read()
	if err != nil {
		return Op{}, err
	}

	if !isBrainfuckChar(c) {
		return this.Next()
	}

	if isRepeatable(c) {
		operand, err := this.countRepeats(c)
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

func (this Lexer) countRepeats(target rune) (uint, error) {
	var count uint = 0

	for {
		next, err := this.peek()
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}

		if next != target {
			break
		}

		if _, err = this.read(); err != nil {
			return 0, err
		}

		count = count + 1
	}

	return count, nil
}

func (this Lexer) peek() (rune, error) {
	c, err := this.read()
	if err != nil {
		return 0, err
	}

	if err := this.unread(); err != nil {
		return 0, fmt.Errorf("unreading after peek: %w", err)
	}

	return c, nil
}

func (this Lexer) read() (rune, error) {
	c, _, err := this.reader.ReadRune()
	if err != nil {
		return 0, err
	}
	return c, nil
}

func (this Lexer) unread() error {
	return this.reader.UnreadRune()
}
