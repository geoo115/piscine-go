package piscine

func ToLower(s string) string {
	var lowerCase string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			lowerCase += string(char + 32)
		} else {
			lowerCase += string(char)
		}
	}
	return lowerCase
}
