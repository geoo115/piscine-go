package piscine

func RecursiveFactorial(nb int) int {
	limit := 50
	if nb < 0 || nb > limit {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
