package bfgo

import (
	"errors"
	"fmt"
)

var (
	ErrEOF              = errors.New("end of file")
	ErrInput            = errors.New("input error")
	ErrOutput           = errors.New("output error")
	ErrPointerUnderflow = errors.New("pointer underflow")
	ErrPointerOverflow  = errors.New("pointer overflow")
)

func NewPointerUnderflowError(pc int) error {
	return fmt.Errorf("%w at pc=%d", ErrPointerUnderflow, pc)
}

func NewPointerOverflowError(pc int) error {
	return fmt.Errorf("%w at pc=%d", ErrPointerOverflow, pc)
}

func NewInputError(err error) error {
	return fmt.Errorf("%w: %w", ErrInput, err)
}

func NewOutputError(err error) error {
	return fmt.Errorf("%w: %w", ErrOutput, err)
}
