package main

import "math"

func maxValue(nums []int) []int {
	n := len(nums)

	suffixMin := make([]int, n+1)
	suffixMin[n] = math.MaxInt

	for i := n - 1; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	result := make([]int, n)

	start := 0

	for start < n {
		current := start
		componentMax := nums[start]

		for current+1 < n && componentMax > suffixMin[current+1] {
			current++

			componentMax = max(componentMax, nums[current])
		}

		for i := start; i <= current; i++ {
			result[i] = componentMax
		}

		start = current + 1
	}

	return result
}
