package piscine

func AlphaCount(s string) int {
	count := 0
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count++
		} else if char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}
