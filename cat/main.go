package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, file := range args {
			content, err := os.ReadFile(file)
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				printStr(string(content))
			}
		}
	}
}
