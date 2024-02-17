package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for first := ('0' - 48); first <= 98; first++ {
		for second := first + 1; second <= 99; second++ {
			z01.PrintRune((first / 10) + '0')
			z01.PrintRune((first % 10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune((second / 10) + '0')
			z01.PrintRune((second % 10) + '0')
			if first != 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
