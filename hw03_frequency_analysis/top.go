package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var splitWordsRegexp = regexp.MustCompile(`\s-\s|[!"#$%&'()*+,\\./:;<=>?@[\]^_\x60{|}~\s]+`)

func Top10(text string) []string {
	wordFrequency := make(map[string]int)

	text = strings.ToLower(text)
	for _, word := range splitWordsRegexp.Split(text, -1) {
		if word != "" {
			wordFrequency[word]++
		}
	}

	words := make([]string, 0, len(wordFrequency))
	for word := range wordFrequency {
		words = append(words, word)
	}

	sort.SliceStable(words, func(i, j int) bool {
		if wordFrequency[words[i]] == wordFrequency[words[j]] {
			return words[i] < words[j]
		}

		return wordFrequency[words[i]] > wordFrequency[words[j]]
	})

	bound := int(math.Min(10, float64(len(words))))
	return words[:bound]
}
