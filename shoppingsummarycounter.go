package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	cart := make(map[string]int)
	var items string
	for _, e := range str {
		if e == 32 {
			cart[items] += 1
			items = ""
		} else if e != 32 {
			items += string(byte(e))
		}
	}
	cart[items] += 1

	return cart
}
