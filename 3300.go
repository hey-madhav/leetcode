package main

import (
	"math"
)

func minElement(nums []int) int {
	var result = math.MaxInt

	for _, num := range nums {
		count := 0
		for num > 0 {
			count += num % 10
			num /= 10
		}

		result = min(count, result)
	}

	return result
}
