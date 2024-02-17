package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	sl := []string{}
	count := 0
	for _, char := range s {
		if char != ' ' || (char <= 9 && char >= 13) {
			str += string(char)
			count++

		} else if count != 0 {
			sl = append(sl, str)
			str = ""
			count = 0
		}
	}
	if str != "" {
		sl = append(sl, str)
	}
	return sl
}
