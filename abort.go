package piscine

func Abort(a, b, c, d, e int) int {
	s := []int{a, b, c, d, e}
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ { // use j as loop variable
			if s[j] < s[i] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s[2]
}
