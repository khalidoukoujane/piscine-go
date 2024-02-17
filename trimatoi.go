package piscine

func TrimAtoi(s string) int {
	isNegative := false
	res := 0
	sign := 1
	isEmpty := true
	for _, char := range s {
		if isEmpty && !isNegative && char == '-' {
			isNegative = true
			sign = -1
		} else if char >= '0' && char <= '9' {
			res *= 10
			res += int(char - '0')
			isEmpty = false
		}
	}
	return res * sign
}
