package main

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	var result = make([]int, len(wordsQuery))
	for i, query := range wordsQuery {
		maxSuffix := 0
		temp := ""
		for j, container := range wordsContainer {
			count := maxSuffixFunc(query, container)
			if count > maxSuffix {
				result[i] = j
				maxSuffix = count
				temp = container
			} else if count == maxSuffix {
				if len(container) < len(temp) {
					result[i] = j
					temp = container
				}
			}
		}

		if maxSuffix == 0 {
			result[i] = smallest(wordsContainer)
		}
	}

	return result
}

func maxSuffixFunc(s, t string) int {
	i := 1

	for (i <= min(len(s), len(t))) && (s[len(s)-i] == t[len(t)-i]) {
		i++
	}

	return i - 1
}

func smallest(s []string) int {
	start := 0
	for i := 1; i < len(s); i++ {
		if len(s[i]) < len(s[start]) {
			start = i
		}
	}

	return start
}
