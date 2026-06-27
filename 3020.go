package main

func maximumLength(nums []int) int {
	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num]++
	}

	var maxLength int
	maxLength = hashMap[1]
	if maxLength%2 == 0 {
		maxLength--
	}

	delete(hashMap, 1)

	for num := range hashMap {
		result := 0
		x := num

		for hashMap[x] > 1 {
			result += 2
			x *= x
		}

		if _, ok := hashMap[x]; ok {
			maxLength = max(maxLength, result+1)
		} else {
			maxLength = max(maxLength, result-1)
		}
	}

	return maxLength
}
