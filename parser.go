package bfgo

import (
	"bufio"
	"fmt"
)

func Parse(reader *bufio.Reader) ([]Op, error) {
	ops := make([]Op, 0, 256)
	stack := make([]int, 0, 32)

	lexer := NewLexer(reader)

	for {
		op, err := lexer.Next()
		if err != nil {
			if err == ErrEOF {
				break
			}
			return nil, err
		}

		switch op.kind {
		case OpJumpFwd:
			stack = append(stack, len(ops))
			ops = append(ops, op)

		case OpJumpBack:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unmatched ']' at operation %d", len(ops))
			}

			// pop from stack
			openPos := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ops[openPos].operand = uint(len(ops))

			jumpBackOp, ok := NewOpWithOperand(']', uint(openPos))
			if !ok {
				return nil, fmt.Errorf("failed to create jump back operation at position %d", len(ops))
			}

			ops = append(ops, jumpBackOp)

		default:
			ops = append(ops, op)
		}
	}

	if len(stack) != 0 {
		return nil, fmt.Errorf("unmatched '[' at operation %d", stack[len(stack)-1])
	}

	return ops, nil
}
