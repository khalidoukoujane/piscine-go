package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	croissant := true
	decroissant := true
	for i := 1; i < len(a); i++ {
		if f(a[i], a[i-1]) < 0 {
			croissant = false
		}
		if f(a[i], a[i-1]) > 0 {
			decroissant = false
		}
	}
	return croissant || decroissant
}
