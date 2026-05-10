package main

import (
	"strings"
	"unicode"
)

func countWordOccurrences(chunks []string, queries []string) []int {
	var sentence string
	for i := range chunks {
		sentence += chunks[i]
	}

	words := sanatisedWords(sentence)

	wordMap := make(map[string]int)

	for i := range words {
		wordMap[words[i]]++
	}

	var result = make([]int, len(queries))
	for i, query := range queries {
		if count, ok := wordMap[query]; ok {
			result[i] = count
		}
	}

	return result

}

func sanatisedWords(sentence string) []string {
	words := strings.Fields(sentence)
	var result []string

	for _, word := range words {
		if strings.Contains(word, "-") {
			parts := strings.Split(word, "-")

			for _, part := range parts {
				if isAllLowercase(part) {
					result = append(result, part)
				}
			}

			if isAllLowercase(word) {
				result = append(result, word)
			}
		} else {
			if isAllLowercase(word) {
				result = append(result, word)
			}
		}
	}

	return result
}

func isAllLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return false
		}
	}
	return true
}
