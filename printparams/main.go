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
	arg := os.Args[1:]
	for _, args := range arg {
		PrintStr(args)
		z01.PrintRune('\n')
	}
}
