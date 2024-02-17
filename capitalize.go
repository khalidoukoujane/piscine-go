package piscine

func lower(s string) string {
	res := []rune(s)
	for i := range s {
		if res[i] >= 'A' && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
	}
	return string(res)
}

func Capitalize(s string) string {
	inWord := false
	res := []rune(lower(s))
	for i := range s {
		if res[i] >= 'a' && res[i] <= 'z' {

			if !inWord {
				res[i] = res[i] - 32
			}
			inWord = true
		} else if res[i] >= '0' && res[i] <= '9' {
			inWord = true
		} else if res[i] >= 'A' && res[i] <= 'Z' {
			inWord = true
		} else {
			inWord = false
		}
	}
	return string(res)
}
