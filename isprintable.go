package piscine

func IsPrintable(s string) bool {
	isprintable := false
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] >= 32 && str[i] <= 126 {
			isprintable = true
		} else {
			return false
		}
	}
	return isprintable
}
