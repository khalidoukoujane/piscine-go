package piscine

func IsUpper(s string) bool {
	isUpper := false
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			isUpper = true
		} else {
			return false
		}
	}
	return isUpper
}
