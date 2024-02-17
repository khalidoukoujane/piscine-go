package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	canAddspace := false
	count := 0
	res := ""
	for i := 0; i < len(str); {
		for i < len(str) && str[i] == ' ' {
			i++
			continue
		}
		if canAddspace {
			res += " "
		} else {
			canAddspace = true
		}
		word := ""
		for i < len(str) && count < 5 {
			if str[i] == ' ' {
				i++
				continue
			}
			if i >= len(str) {
				break
			}
			if str[i] != ' ' {
				count++
				word += string(str[i])
			}
			i++
		}
		count = 0
		if word != "" {
			res += word
		} else {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		}

		word = ""
		i++
	}

	return res + "\n"
}
