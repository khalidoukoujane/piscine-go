package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i, str := range strs {
		res += str
		if i != len(strs)-1 {
			res += sep
		}
	}
	return res
}
