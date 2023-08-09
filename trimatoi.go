package piscine

func TrimAtoi(s string) int {
	// Initialize variables
	numberFound := false
	numberCount := 0
	result := 0
	isPositive := true
	// Iterate over each character in the string
	for _, char := range s {
		// Check if the character is a digit
		if '0' <= char && char <= '9' {
			// If the first digit encountered is a 0, skip it
			if char == '0' && !numberFound {
				continue
			}
			// Add the digit to the result
			result = result*10 + int(char-'0')
			// Set the numberFound flag and increment the numberCount variable
			numberFound = true
			numberCount++
		} else if char == '-' && !numberFound {
			// If a minus sign is encountered before any digits, set the isPositive flag to false
			isPositive = false
		}
	}
	// Negate the result if isPositive is false
	if !isPositive {
		result *= -1
	}
	return result
}
