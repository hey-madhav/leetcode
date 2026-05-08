package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	maxProduct, minProduct, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = max(nums[i], maxProduct*nums[i])
		minProduct = min(nums[i], minProduct*nums[i])

		result = max(result, maxProduct)
	}

	return result
}
