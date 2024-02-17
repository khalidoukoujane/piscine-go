package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	res := make([]int, max-min)
	i := 0
	for min < max {
		res[i] = min
		min++
		i++
	}
	return res
}
