package main

func maxTotalValue(nums []int, k int) int64 {
	maxi, mini := nums[0], nums[0]

	for _, num := range nums {
		maxi = max(maxi, num)
		mini = min(mini, num)
	}

	sum := maxi - mini

	return int64(k * sum)
}
