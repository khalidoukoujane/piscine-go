package piscine

func IntToStr(n int) string {
	res := ""
	if n < 0 {
		return ""
	}
	if n == 0 {
		res += string("0")
	}
	for n > 0 {
		res += string(rune((n % 10) + '0'))
		n /= 10
	}
	return res
}

func PrintNbrInOrder(n int) {
	str := IntToStr(n)
	result := []rune(str)
	i := 1
	for i < len(result) {
		j := i
		for j > 0 {
			if result[j-1] > result[j] {
				result[j], result[j-1] = result[j-1], result[j]
			}
			j--
		}
		i++
	}
	PrintStr(string(result))
}
