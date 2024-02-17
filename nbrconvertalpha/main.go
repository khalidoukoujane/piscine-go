package main

import (
	"os"

	"github.com/01-edu/z01"
)

func ToUpper(s string) string {
	res := []rune(s)
	for i := range s {
		if res[i] >= 'a' && res[i] <= 'z' {
			res[i] = res[i] - 32
		}
	}
	return string(res)
}

func MiniAtoi(s string) int {
	res := 0
	for _, char := range s {
		res *= 10
		if char >= '0' && char <= '9' {
			res += int(char - '0')
		}
	}
	return res
}

func PrintStr(s string) {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		z01.PrintRune(str[i])
	}
}

func main() {
	args := os.Args
	Latin := "abcdefghijklmnopqrstuvwxyz"
	sliceOfalpha := []rune(Latin)
	if len(args) < 2 {
		return
	}
	for i := 1; i < len(args); i++ {
		index := MiniAtoi(args[i])
		if len(args) > 2 && args[1] == "--upper" {
			if index >= 1 && index <= 26 {
				PrintStr(ToUpper(string(sliceOfalpha[index-1])))
			} else if index == 0 {
				PrintStr("")
			} else {
				z01.PrintRune(' ')
			}
		} else {
			if index >= 1 && index <= 26 {
				PrintStr(string(sliceOfalpha[index-1]))
			} else {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
