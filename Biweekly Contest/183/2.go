package main

import "math"

func minOperations(nums []int, k int) int {
	evenCost := make([]int, k)
	oddCost := make([]int, k)

	for i := 0; i < k; i++ {
		costEven := 0
		costOdd := 0

		for j := 0; j < len(nums); j++ {
			remainder := nums[j] % k
			diff := abs(remainder - i)
			steps := min(diff, k-diff)

			if j%2 == 0 {
				costEven += steps
			} else {
				costOdd += steps
			}
		}

		evenCost[i] = costEven
		oddCost[i] = costOdd
	}

	result := math.MaxInt

	for x := 0; x < k; x++ {
		for y := 0; y < k; y++ {
			if x == y {
				continue
			}

			result = min(result, evenCost[x] + oddCost[y])
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
