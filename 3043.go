package main

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixMap := make(map[int]bool)

	for _, val := range arr1 {
		for val > 0 {
			prefixMap[val] = true
			val /= 10
		}
	}

	longestPrefix := 0

	for _, val := range arr2 {
		for val > 0 {
			if _, ok := prefixMap[val]; ok {
				longestPrefix = max(longestPrefix, len(strconv.Itoa(val)))
			}
			val /= 10
		}
	}

	return longestPrefix
}
