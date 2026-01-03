package bfgo

func isBrainfuckChar(c rune) bool {
	switch c {
	case '+', '-', '<', '>', '.', ',', '[', ']':
		return true
	default:
		return false
	}
}

func isRepeatable(c rune) bool {
	switch c {
	case '+', '-', '<', '>', '.':
		return true
	default:
		return false
	}
}
