package piscine

func Split(s, sep string) []string {
	var res []string
	var str string
	for i := 0; i < len(s); i++ {
		if i+len(sep) > len(s) || string(s[i:i+len(sep)]) != sep {
			str += string(s[i])
		} else {
			res = append(res, str)
			str = ""
			i += len(sep) - 1
		}
	}
	if str != "" {
		res = append(res, str)
	}
	return res
}
