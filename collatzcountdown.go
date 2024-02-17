package piscine

func CollatzCountdown(start int) int {
	if start <= 0 {
		return -1
	}
	count := 0
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			count++
		} else {
			start = 3*start + 1
			count++
		}
	}
	return count
}
