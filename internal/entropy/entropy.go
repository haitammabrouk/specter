package entropy

import (
	"math"
	"strings"
)

type TokenEntropy struct {
	Token   string
	Entropy float64
}

func shannonEntropy(word string) float64 {
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

func ComputeEntropyPerToken(line string) []TokenEntropy {
	tokens := strings.Fields(line)
	entropyDic := make([]TokenEntropy, 0, len(tokens))

	for _, token := range tokens {
		entropyDic = append(entropyDic, TokenEntropy{
			Token:   token,
			Entropy: shannonEntropy(token),
		})
	}

	return entropyDic
}
