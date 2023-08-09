package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		switch {
		case char >= 'a' && char <= 'z':
			result += string((char-'a'+14)%26 + 'a')
		case char >= 'A' && char <= 'Z':
			result += string((char-'A'+14)%26 + 'A')
		default:
			result += string(char)
		}
	}
	return result
}
