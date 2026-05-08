package main

func findValidElements(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}

	n := len(nums)
	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i+1] > rightMax[i+1] {
			rightMax[i] = nums[i+1]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}

	result := make([]int, 0, n)
	result = append(result, nums[0])
	leftMax := nums[0]

	for i := 1; i < n-1; i++ {
		if nums[i] > leftMax || nums[i] > rightMax[i] {
			result = append(result, nums[i])
		}
		if nums[i] > leftMax {
			leftMax = nums[i]
		}
	}

	result = append(result, nums[n-1])
	return result
}
