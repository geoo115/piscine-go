package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if table[j] < table[i] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}
