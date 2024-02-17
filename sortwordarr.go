package piscine

func compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		j := i
		for j > 0 {
			if compare(a[j-1], a[j]) == 1 {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j--
		}
	}
}
