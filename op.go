package bfgo

import "fmt"

type OpKind uint8

const (
	OpInc OpKind = iota
	OpDec
	OpLeft
	OpRight
	OpOutput
	OpInput
	OpJumpFwd
	OpJumpBack
)

type Op struct {
	kind    OpKind
	operand uint
}

func NewOp(c rune) (Op, bool) {
	switch c {
	case '+':
		return Op{kind: OpInc, operand: 1}, true
	case '-':
		return Op{kind: OpDec, operand: 1}, true
	case '<':
		return Op{kind: OpLeft, operand: 1}, true
	case '>':
		return Op{kind: OpRight, operand: 1}, true
	case '.':
		return Op{kind: OpOutput, operand: 0}, true
	case ',':
		return Op{kind: OpInput, operand: 0}, true
	case '[':
		return Op{kind: OpJumpFwd, operand: 0}, true
	case ']':
		return Op{kind: OpJumpBack, operand: 0}, true
	default:
		return Op{}, false
	}
}

func NewOpWithOperand(c rune, operand uint) (Op, bool) {
	switch c {
	case '+':
		return Op{kind: OpInc, operand: operand}, true
	case '-':
		return Op{kind: OpDec, operand: operand}, true
	case '<':
		return Op{kind: OpLeft, operand: operand}, true
	case '>':
		return Op{kind: OpRight, operand: operand}, true
	case '.':
		return Op{kind: OpOutput, operand: operand}, true
	case '[':
		return Op{kind: OpJumpFwd, operand: operand}, true
	case ']':
		return Op{kind: OpJumpBack, operand: operand}, true
	default:
		return Op{}, false
	}
}

func (op Op) String() string {
	switch op.kind {
	case OpInc:
		return fmt.Sprintf("OpInc(%d)", op.operand)
	case OpDec:
		return fmt.Sprintf("OpDec(%d)", op.operand)
	case OpLeft:
		return fmt.Sprintf("OpLeft(%d)", op.operand)
	case OpRight:
		return fmt.Sprintf("OpRight(%d)", op.operand)
	case OpOutput:
		return "OpOutput"
	case OpInput:
		return "OpInput"
	case OpJumpFwd:
		return fmt.Sprintf("OpJumpFwd(%d)", op.operand)
	case OpJumpBack:
		return fmt.Sprintf("OpJumpBack(%d)", op.operand)
	default:
		return "UnknownOp"
	}
}
