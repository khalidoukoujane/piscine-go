package piscine

func IsLower(s string) bool {
	isUpper := false
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			isUpper = true
		} else {
			return false
		}
	}
	return isUpper
}
