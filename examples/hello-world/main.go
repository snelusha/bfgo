package main

import (
	"bufio"
	"os"

	"bfgo"
)

func main() {
	source, err := os.Open("hw.bf")
	if err != nil {
		panic(err)
	}
	defer source.Close()

	interpreter, err := bfgo.NewInterpreter(bufio.NewReader(source), os.Stdout, os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := interpreter.Execute(); err != nil {
		panic(err)
	}
}
