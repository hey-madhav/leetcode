package main

func countMajoritySubarrays(nums []int, target int) int {
	n := len(nums)
	ans := 0

	for left := 0; left < n; left++ {
		countTarget := 0

		for right := left; right < n; right++ {

			if nums[right] == target {
				countTarget++
			}

			length := right - left + 1

			if 2*countTarget > length {
				ans++
			}
		}
	}

	return ans
}
