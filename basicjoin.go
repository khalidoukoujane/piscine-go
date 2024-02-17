package piscine

func BasicJoin(elems []string) string {
	res := ""
	for _, elem := range elems {
		res += elem
	}
	return res
}
