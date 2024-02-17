package piscine

func Map(f func(int) bool, a []int) []bool {
	res := []bool{}
	for i := 0; i < len(a); i++ {
		res = append(res, f(a[i]))
	}
	return res
}
