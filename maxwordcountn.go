package piscine

import (
	"sort"
	"strings"
)

func MaxWordCountN(text string, n int) map[string]int {
	// Split the string into words
	words := strings.Split(text, " ")

	// Count the occurrences of each word
	// Count the occurrences of each word
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	// Sort the words by occurrence count (descending) and then by ASCII value (ascending)
	type wordCount struct {
		word  string
		count int
	}
	sorted := make([]wordCount, 0, len(counts))
	for word, count := range counts {
		sorted = append(sorted, wordCount{word, count})
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count == sorted[j].count {
			return sorted[i].word < sorted[j].word
		}
		return sorted[i].count > sorted[j].count
	})

	// Build the result map with the top n words
	result := make(map[string]int)
	for i := 0; i < n && i < len(sorted); i++ {
		result[sorted[i].word] = sorted[i].count
	}
	return result
}
