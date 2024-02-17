package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a
	*a /= *b
	*b = tmp % *b
}
