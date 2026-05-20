package main

func findThePrefixCommonArray(A []int, B []int) []int {
	var prefixArray []int
	var current = 0
	var dictionary = make(map[int]bool)

	for current < len(A) {
		var commonPrefixCount = 0
		dictionary[A[current]] = true

		for i := 0; i <= current; i++ {
			if dictionary[B[i]] {
				commonPrefixCount++
			}
		}

		prefixArray = append(prefixArray, commonPrefixCount)
		current++
	}

	return prefixArray
}
