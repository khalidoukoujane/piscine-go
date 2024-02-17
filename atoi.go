package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	for i := range s {
		result *= 10
		if i == 0 && s[i] == '-' {
			sign = -1
		} else if i == 0 && s[i] == '+' {
			sign = 1
		} else if s[i] > '9' || s[i] < '0' {
			result = 0
			break
		} else {
			result += int(s[i] - '0')
		}
	}
	return result * sign
}
