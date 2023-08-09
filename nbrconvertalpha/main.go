package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	caps := false
	if len(arguments) > 0 && arguments[0] == "--upper" {
		caps = true
		arguments = arguments[1:]
	}
	for _, arg := range arguments {
		numv := 0
		for _, v := range arg {
			if v >= '0' && v <= '9' {
				numv = numv*10 + int(v-'0')
			}
		}
		if numv >= 1 && numv <= 26 {
			if caps == false {
				z01.PrintRune(rune('a' + numv - 1))
			} else {
				z01.PrintRune(rune('A' + numv - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
