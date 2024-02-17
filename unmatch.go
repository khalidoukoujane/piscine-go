package piscine

func Unmatch(a []int) int {
	for _, num := range a {
		rep := 0
		for _, pair := range a {
			if num == pair {
				rep++
			}
		}
		if rep == 1 || rep%2 == 1 {
			return num
		}
	}
	return -1
}
