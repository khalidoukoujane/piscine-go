package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if n <= 0 || n > len(str) {
		return 0
	}
	return str[n-1]
}
