package main

func limitOccurrences(nums []int, k int) []int {
	count := k - 1
	var result []int

	result = append(result, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if count > 0 {
				result = append(result, nums[i])
				count--
			}
		} else {
			result = append(result, nums[i])
			count = k - 1
		}
	}

	return result
}
