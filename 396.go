package main

import "math"

func maxRotateFunction(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	rotatedValue, totalSum := rotateFunction(nums)
	var result = math.MinInt
	for i := 0; i < len(nums); i++ {
		rotatedValue = rotatedValue + totalSum - len(nums)*nums[len(nums)-i]
		result = max(result, rotatedValue)
	}

	return result

}

func rotateFunction(nums []int) (int, int) {
	sum := 0
	totalSum := 0

	for i := range nums {
		sum += i * nums[i]
		totalSum += nums[i]
	}

	return sum, totalSum
}
