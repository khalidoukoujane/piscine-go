package piscine

func AppendRange(min, max int) []int {
	res := []int{}
	if min >= max {
		res = nil
		return res
	}
	for min < max {
		res = append(res, min)
		min++
	}
	return res
}
