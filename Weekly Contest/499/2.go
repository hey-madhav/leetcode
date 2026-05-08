package main

import "sort"

type vowelInfo struct {
	char  byte
	freq  int
	first int
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

func sortVowels(s string) string {
	if len(s) <= 2 {
		return s
	}

	var frequencyMap = make(map[byte]int)
	var firstOccurrence = make(map[byte]int)

	for i := range s {
		if isVowel(s[i]) {
			frequencyMap[s[i]]++
			if _, ok := firstOccurrence[s[i]]; !ok {
				firstOccurrence[s[i]] = i
			}
		}
	}

	var vowels []vowelInfo
	for char, freq := range frequencyMap {
		vowels = append(vowels, vowelInfo{char, freq, firstOccurrence[char]})
	}

	sort.Slice(vowels, func(i, j int) bool {
		if vowels[i].freq != vowels[j].freq {
			return vowels[i].freq > vowels[j].freq
		}
		return vowels[i].first < vowels[j].first
	})

	var sortedVowels []byte
	for _, v := range vowels {
		for k := 0; k < v.freq; k++ {
			sortedVowels = append(sortedVowels, v.char)
		}
	}

	var result []byte
	var idx = 0
	for i := range s {
		char := s[i]
		if isVowel(char) {
			result = append(result, sortedVowels[idx])
			idx++
		} else {
			result = append(result, char)
		}
	}

	return string(result)
}
