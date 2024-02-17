package main

import (
	"fmt"
	"os"
)

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		if s[i] >= '0' && s[i] <= '9' {
			res += int(s[i] - '0')
		}
	}
	return res
}

func display(index int, res string) {
	for i := len(res) - index; i < len(res); i++ {
		fmt.Printf("%c", rune(res[i]))
	}
	// fmt.Printf("\n")
}

func contains(s string, toFind string) int {
	l := 0
	for _, char := range s {
		for _, c := range toFind {
			if char == c {
				l++
				break
			}
		}
	}
	return l
}

func main() {
	args := os.Args
	if args[1] != "-c" || len(args) < 3 {
		return
	}
	index := 0
	valid := contains(args[2], "0123456789")
	if len(args[2]) == valid {
		index = BasicAtoi(args[2])
	}
	for i := 3; i < len(args); i++ {
		content, err := os.ReadFile(args[i])
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		res := string(content)
		if len(args[3:]) > 1 {
			if i > 3 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", args[i])
		}
		display(index, res)
	}
}
