package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // ignore the first argument, which is the program name
	for i := len(arguments) - 1; i >= 0; i-- {
		arg := arguments[i]
		for _, ch := range arg {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
