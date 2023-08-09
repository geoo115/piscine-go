package piscine

func IsUpper(s string) bool {
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			continue
		}
		if ch < 'A' || ch > 'Z' {
			return false
		}
	}
	return true
}
