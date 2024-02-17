package piscine

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		if s[i] >= '0' && s[i] <= '9' {
			res += int(s[i] - '0')
		}
	}
	return res
}
