package main

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	var left, result = 0, math.MaxInt
	
	var sum = 0
	for right, num := range nums {
		sum += num

		for sum >= target {
			result = min(result, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		return 0
	}

	return result
}
