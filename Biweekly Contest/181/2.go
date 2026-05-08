package main

func compareBitonicSums(nums []int) int {
	var peakIndex int

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			peakIndex = i - 1
			break
		}
	}

	firstSum, secondSum := sum(nums[:peakIndex+1]), sum(nums[peakIndex:])

	if firstSum > secondSum {
		return 0
	}

	if secondSum > firstSum {
		return 1
	}

	return -1
}

func sum(nums []int) int {
	var sum = 0
	for i := range nums {
		sum += nums[i]
	}

	return sum
}
