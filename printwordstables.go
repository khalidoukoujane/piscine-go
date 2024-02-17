package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		PrintStr(s)
		z01.PrintRune('\n')
	}
}
