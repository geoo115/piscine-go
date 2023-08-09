package piscine

func Map(f func(int) bool, a []int) []bool {
	r := []bool{}
	for _, s := range a {
		r = append(r, f(s))
	}
	return r
}

// package piscine
// func Map(f func(int) bool, arr []int) []bool {
// 	length := 0
// 	for l := range arr {
// 		length = l + 1
// 	}
// 	boolAr := make([]bool, length)
// 	for i := range arr {
// 		boolAr[i] = f(arr[i])
// 	}
// 	return boolAr
// }
