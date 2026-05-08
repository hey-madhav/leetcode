package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	var anagramMap = make(map[string][]string)

	for ind := range strs {
		slice := []byte(strs[ind])
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})

		anagramMap[string(slice)] = append(anagramMap[string(slice)], strs[ind])
	}

	var results = make([][]string, 0, len(anagramMap))
	for _, anagram := range anagramMap {
		results = append(results, anagram)
	}

	return results
}
