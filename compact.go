package piscine

func Compact(ptr *[]string) int {
	CompactedArray := []string{}

	for _, elem := range *ptr {
		if elem != "" {
			CompactedArray = append(CompactedArray, elem)
		}
	}

	*ptr = CompactedArray
	return len(CompactedArray)
}
