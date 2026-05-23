package main

func minimumSwaps(nums []int) int {
	nonZeroIndex := len(nums) - 1
	zeroIndex := 0

	var count = 0
	for zeroIndex < nonZeroIndex {
		for nonZeroIndex >= 0 && nums[nonZeroIndex] == 0 {
			nonZeroIndex--
		}

		for zeroIndex < len(nums) && nums[zeroIndex] != 0 {
			zeroIndex++
		}

		if zeroIndex < nonZeroIndex {
			nums[zeroIndex], nums[nonZeroIndex] = nums[nonZeroIndex], nums[zeroIndex]
			zeroIndex++
			nonZeroIndex--
			count++
		}
	}

	return count
}
