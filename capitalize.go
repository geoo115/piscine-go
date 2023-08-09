package piscine

func prim(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	ar := []rune(s)
	letra := true
	for i := 0; i < len(s); i++ {
		if prim(ar[i]) == true && letra {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = 'A' - 'a' + ar[i]
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = 'a' - 'A' + ar[i]
		} else if prim(ar[i]) == false {
			letra = true
		}
	}
	return string(ar)
}

// //package piscine
// func check(a rune) bool {
// 	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
// 		return true
// 	}
// 	return false
// }
// func Capitalize(s string) string {
// 	runes := []rune(s)
// 	first := true
// 	for i := range runes {
// 		if check(runes[i]) == true && first {
// 			if runes[i] >= 'a' && runes[i] <= 'z' {
// 				runes[i] -= 32
// 			}
// 			first = false
// 		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
// 			runes[i] += 32
// 		} else if check(runes[i]) == false {
// 			first = true
// 		}
// 	}
// 	return string(runes)
// }
