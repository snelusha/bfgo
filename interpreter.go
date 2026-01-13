package bfgo

import (
	"bufio"
	"fmt"
	"io"
)

type Interpreter struct {
	memory [30000]byte
	ptr    int
	pc     int
	ops    []Op
	output io.Writer
	input  io.Reader
}

func NewInterpreter(reader *bufio.Reader, output io.Writer, input io.Reader) (*Interpreter, error) {
	ops, err := Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	return &Interpreter{
		memory: [30000]byte{},
		ptr:    0,
		pc:     0,
		ops:    ops,
		output: output,
		input:  input,
	}, nil
}

func (i *Interpreter) Execute() error {
	for i.pc < len(i.ops) {
		op := i.ops[i.pc]

		switch op.kind {
		case OpInc:
			i.memory[i.ptr] += byte(op.operand)
		case OpDec:
			i.memory[i.ptr] -= byte(op.operand)
		case OpLeft:
			i.ptr -= int(op.operand)
			if i.ptr < 0 {
				return fmt.Errorf("pointer underflow at pc=%d", i.pc)
			}
		case OpRight:
			i.ptr += int(op.operand)
			if i.ptr >= len(i.memory) {
				return fmt.Errorf("pointer overflow at pc=%d", i.pc)
			}
		case OpOutput:
			for range op.operand {
				if _, err := i.output.Write([]byte{i.memory[i.ptr]}); err != nil {
					return fmt.Errorf("output error: %w", err)
				}
			}
		case OpInput:
			buf := make([]byte, 1)
			if _, err := i.input.Read(buf); err != nil {
				return fmt.Errorf("input error: %w", err)
			}
			i.memory[i.ptr] = buf[0]
		case OpJumpFwd:
			if i.memory[i.ptr] == 0 {
				i.pc = int(op.operand)
				continue
			}
		case OpJumpBack:
			if i.memory[i.ptr] != 0 {
				i.pc = int(op.operand)
				continue
			}
		}

		i.pc++
	}
	return nil
}
