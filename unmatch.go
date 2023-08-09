package piscine

func Unmatch(a []int) int {
	count := 0
	for _, v1 := range a {
		for _, v2 := range a {
			if v1 == v2 {
				count++
			}
		}
		if count%2 > 0 {
			return v1
		}
	}
	return -1
}
