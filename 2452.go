package main

func twoEditWords(queries []string, dictionary []string) []string {
	var result []string
	for _, query := range queries {
		for _, dict := range dictionary {
			if findDifferentCharacterCount(query, dict) <= 2 {
				result = append(result, query)
			}
		}
	}

	return result
}

func findDifferentCharacterCount(str1, str2 string) int {
	var diffCount int
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			diffCount++
		}
	}

	return diffCount
}
