package main

func leftRightDifference(nums []int) []int {
    rightSum := 0

	for i := range nums {
		rightSum += nums[i]
	}

	leftSum := 0

	var result = make([]int, len(nums))

	for i := range nums {
		rightSum -= nums[i]
		result[i] = abs(rightSum - leftSum)
		leftSum += nums[i]
	}

	return result
}