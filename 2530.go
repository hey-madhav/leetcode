package main

import (
	"math"
)

func maxKelements(nums []int, k int) int64 {
	var result int64

	for k > 0 {
		maxElement, index := findBiggestElementOfArrayWithIndex(nums)
		result += int64(maxElement)
		nums[index] = int(math.Ceil((float64(maxElement) / 3)))
		k--
	}
	return result
}

func findBiggestElementOfArrayWithIndex(arr []int) (int, int) {
	index := 0
	output := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > output {
			output = arr[i]
			index = i
		}
	}
	return output, index
}
