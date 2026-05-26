package main

func numberOfSpecialChars(word string) int {
	byteMap := make(map[byte]bool)

	for i := range word {
		byteMap[word[i]] = true
	}

	result := 0
	for val := range byteMap {
		if byteMap[val+32] || byteMap[val-32] {
			result++
			delete(byteMap, val)
		}
	}

	return result
}
