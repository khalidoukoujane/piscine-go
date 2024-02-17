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

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func main() {
	args := os.Args
	args = args[1:]
	for i := 0; i < len(args); i++ {
		j := i
		for j > 0 {
			if Compare(args[j-1], args[j]) == 1 {
				args[j-1], args[j] = args[j], args[j-1]
			}
			j--
		}
	}
	for _, arg := range args {
		PrintStr(arg)
		z01.PrintRune('\n')
	}
}
