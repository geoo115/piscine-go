package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sorted := true
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			sorted = false
			break
		}
	}
	if !sorted {
		sorted = true
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) < 0 {
				sorted = false
				break
			}
		}
	}
	return sorted
}
