package piscine

func ConcatParams(args []string) string {
	res := ""

	for i, str := range args {
		res += str
		if i < len(args)-1 {
			res += string("\n")
		}
	}
	return string(res)
}
