package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb == 0 {
		return res
	}
	if nb < 0 {
		return 0
	} else {
		for i := 1; i < nb+1; i++ {
			res *= i
			if res < 0 {
				return 0
			}
		}
		return res
	}
}
