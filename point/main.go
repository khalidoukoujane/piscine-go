package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func show(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func PrintNum(nb int) {
	char := '0'
	if nb/10 > 0 {
		PrintNum(nb / 10)
	}
	for i := 0; i < nb%10; i++ {
		char++
	}
	z01.PrintRune(char)
}

func main() {
	points := &point{}

	setPoint(points)
	show("x = ")
	PrintNum(points.x)
	show(", y = ")
	PrintNum(points.y)
	z01.PrintRune('\n')
}
