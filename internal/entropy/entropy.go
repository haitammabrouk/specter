package entropy

import (
	"math"
)

func ComputeShannonEntropy(word string) float64 {
	freq := make(map[rune]int)
	totalLength := 0
	for _, r := range word {
		freq[r]++
		totalLength++
	}

	entropy := 0.0
	for _, count := range freq {
		p := float64(count) / float64(totalLength)
		entropy -= p * math.Log2(p)
	}
	return entropy
}