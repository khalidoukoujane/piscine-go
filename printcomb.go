package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	first := '0'
	secend := '1'
	third := '2'
	for first <= '7' {
		secend = first + 1
		for secend <= '8' {
			third = secend + 1
			for third <= '9' {
				z01.PrintRune(first)
				z01.PrintRune(secend)
				z01.PrintRune(third)
				if first != '7' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				third++
			}
			secend++
		}
		first++
	}
	z01.PrintRune('\n')
}
