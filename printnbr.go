package piscine

import "github.com/01-edu/z01"

func showstr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		showstr("-9223372036854775808")
		return
	}
	if n < 0 {
		n = -n
		z01.PrintRune('-')
	}
	if n <= 9 {
		z01.PrintRune(rune(n + '0'))
	} else {
		PrintNbr(n / 10)
		z01.PrintRune(rune((n % 10) + '0'))
	}
}
