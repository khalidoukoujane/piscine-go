package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	for i := range s {
		if string(str[i:i+len(toFind)]) == toFind {
			return i
		}
	}
	return -1
}
