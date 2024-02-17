package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
}

func main() {
	args := os.Args[1:]
	l := len(args) - 1
	for l >= 0 {
		PrintStr(args[l])
		z01.PrintRune('\n')
		l--
	}
}
