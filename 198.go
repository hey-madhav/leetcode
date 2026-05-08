package main

func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    if len(nums) == 1 {
        return nums[0]
    }
    
	dp := make([]int, len(nums))
	dp[len(dp)-1] = nums[len(nums)-1]
	dp[len(dp)-2] = max(nums[len(nums)-2], nums[len(nums)-1])

	for i := len(nums) - 3; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}

	return dp[0]
}
