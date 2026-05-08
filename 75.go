package main

func sortColors(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	for right := left; right < len(nums); right++ {
		if nums[right] == 1 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
