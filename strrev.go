package piscine

func StrRev(s string) string {
	var result string
	str := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}
