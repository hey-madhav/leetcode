package main

func concatWithReverse(nums []int) []int {
	var result = make([]int, 2*len(nums))
	copy(result, nums)


	index := 1
	for i := len(nums); i < len(result); i++ {
		result[i] = nums[len(nums)-index]
		index++
	}

	return result
}
