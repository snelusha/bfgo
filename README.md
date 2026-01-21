# bfgo

A minimal Brainfuck interpreter in Go. It lexes, resolves loop jumps during parse, and runs programs on a 30,000-byte tape with basic safety checks.

## Quick start

- Go 1.25+.
- Run the sample: `go run ./examples/hello-world`
- Run tests: `go test ./...`

## As a library

```go
src := bufio.NewReader(strings.NewReader("+++."))
interp, err := bfgo.NewInterpreter(src, os.Stdout, os.Stdin)
if err != nil {
	log.Fatal(err)
}
if err := interp.Execute(); err != nil {
	log.Fatal(err)
}
```

Behavior notes:
- Repeated `+ - < > .` are collapsed with counts for efficiency.
- Unmatched `[`/`]` returns a parse error; pointer under/overflow is guarded.
- Input reads a single byte at the pointer; EOF surfaces as `ErrEOF`.

## Layout

- `lexer.go`, `parser.go`, `interpreter.go`: core pipeline.
- `op.go`, `errors.go`, `helpers.go`: op definitions and helpers.
- `examples/hello-world/`: runnable Brainfuck sample.
