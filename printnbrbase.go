package piscine

func CheckDub(s string) bool {
	dejaVue := make(map[rune]bool)
	for _, char := range s {
		if dejaVue[char] {
			return true
		}
		dejaVue[char] = true
	}
	return false
}

func checkSign(s string) bool {
	for _, char := range s {
		if char == '-' || char == '+' {
			return true
		}
	}
	return false
}

func PrintNbrBase(nbr int, base string) {
	res := ""
	if checkSign(base) {
		PrintStr("NV")
		return
	}
	if CheckDub(base) {
		PrintStr("NV")
		return
	}
	if len(base) < 2 {
		PrintStr("NV")
		return
	}
	if nbr == -9223372036854775808 {
		PrintStr("-")
		PrintNbrBase(9, base)
		PrintNbrBase(223372036854775808, base)
		return
	}
	if nbr < 0 {
		PrintStr("-")
		nbr = -nbr
	}
	for nbr > 0 {
		index := nbr % len(base)
		res = string(base[index]) + res
		nbr /= len(base)
	}
	PrintStr(res)
}
