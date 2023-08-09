package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	baseLen := len(base)
	if baseLen < 2 {
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr >= baseLen {
		PrintNbrBase(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}
