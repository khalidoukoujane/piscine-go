package piscine

func Swap(a *int, b *int) {
	*b, *a = *a, *b
}
