package piscine

func SortIntegerTable(table []int) {
	i := 1
	for i < len(table) {
		j := i
		for j > 0 {
			if table[j-1] > table[j] {
				table[j], table[j-1] = table[j-1], table[j]
			}
			j--
		}
		i++
	}
}
