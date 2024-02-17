package piscine

func ToLower(s string) string {
	res := []rune(s)
	for i := range s {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
	}
	return string(res)
}
