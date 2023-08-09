package piscine

func LastRune(s string) rune {
	ar := []rune(s)
	return ar[len(ar)-1]
}
