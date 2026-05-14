package main

func isGood(nums []int) bool {
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for i := 1; i < len(nums)-1; i++ {
		if countMap[i] != 1 {
			return false
		}
	}

	if countMap[len(nums)-1] != 2 {
		return false
	}

	return true
}
