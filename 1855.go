package main

func maxDistance2(nums1 []int, nums2 []int) int {
	var maxDistance int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] > nums2[j] {
			i++
		} else {
			maxDistance = max(maxDistance, j-i)
			j++
		}
	}

	return maxDistance
}