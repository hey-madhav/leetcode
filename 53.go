package main

func maxSubArray(nums []int) int {
	var result, subArraySum = nums[0], 0

	for i := range nums {
		subArraySum += nums[i]
		result = max(result, subArraySum)
		if subArraySum < 0 {
			subArraySum = 0
		}
	}
	return result
}
