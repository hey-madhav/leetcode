package main

func mapWordWeights(words []string, weights []int) string {
	var result []byte
	for _, word := range words {
		sum := 0

		for _, ch := range word {
			sum += weights[ch-'a']
		}

		result = append(result, byte('z'-sum%len(weights)))
	}

	return string(result)
}
