package piscine

func Compact(ptr *[]string) int {
	// Get the slice from the pointer
	slice := *ptr
	// Count the number of non-empty elements
	count := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			count++
		}
	}
	// Create a new slice with the non-empty elements
	newSlice := make([]string, count)
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			newSlice[j] = slice[i]
			j++
		}
	}
	// Update the pointer to point to the new slice
	*ptr = newSlice
	// Return the number of non-empty elements
	return count
}
