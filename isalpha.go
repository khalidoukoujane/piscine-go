package piscine

func IsAlpha(s string) bool {
	isalpha := false
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			isalpha = true
		} else if char >= 'A' && char <= 'Z' {
			isalpha = true
		} else if char >= '0' && char <= '9' {
			isalpha = true
		} else {
			return false
		}
	}
	return isalpha
}
