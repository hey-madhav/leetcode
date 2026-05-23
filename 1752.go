package main

func check(nums []int) bool {
	length := len(nums) - 1
	if nums[0] < nums[length] {
		for i := 1; i <= length; i++ {
			if nums[i] < nums[i-1] {
				return false
			}
		}
	} else {
		start := 1

		for start <= length {
			if nums[start] < nums[start-1] {
				break
			}
			start++
		}

		for i := start + 1; i <= length; i++ {
			if nums[i] < nums[i-1] {
				return false
			}
		}
	}

	return true
}
