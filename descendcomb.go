package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for first := 99; first >= 1; first-- {
		for second := first - 1; second >= 0; second-- {
			z01.PrintRune(rune(first/10) + '0')
			z01.PrintRune(rune(first%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(second/10) + '0')
			z01.PrintRune(rune(second%10) + '0')
			if first != 1 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
