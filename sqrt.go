package piscine

func Sqrt(nb int) int {
	res := 0
	for i := 1; i <= nb/i; i++ {
		if i*i == nb {
			res = i
		}
	}
	return res
}
