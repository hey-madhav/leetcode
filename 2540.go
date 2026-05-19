package main 

func getCommon(nums1 []int, nums2 []int) int {
	seenMap := make(map[int]bool)

	for _, item := range nums1 {
		seenMap[item] = true
	}

	for _, item := range nums2 {
		if seenMap[item] {
			return item
		}
	}

	return -1
}