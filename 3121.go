package main

func numberOfSpecialChars2(word string) int {
	indexMap := make(map[byte][]int)

	for i := range word {
		indexMap[word[i]] = append(indexMap[word[i]], i)
	}

	var result int

	for ch, ind := range indexMap {
		if indexMap[ch-32] != nil {
			if ind[len(ind)-1] < indexMap[ch-32][0] {
				result++
				delete(indexMap, ch)
			}
		}
	}

	return result
}
