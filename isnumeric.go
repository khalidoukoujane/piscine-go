package piscine

func IsNumeric(s string) bool {
	isnum := false
	for _, char := range s {
		if char >= '0' && char <= '9' {
			isnum = true
		} else {
			return false
		}
	}
	return isnum
}
