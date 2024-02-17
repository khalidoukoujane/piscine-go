package piscine

func ToUpper(s string) string {
	res := []rune(s)
	for i := range s {
		if res[i] >= 'a' && res[i] <= 'z' {
			res[i] = res[i] - 32
		}
	}
	return string(res)
}
