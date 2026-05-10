package main

func maximumJumps(nums []int, target int) int {
	dp := make([]int, len(nums))
    n := len(nums)

    for i := range dp {
        dp[i] = -1
    }

    dp[0] = 0
	for i := 0; i<n ;i++{
        for j:=0;j<i;j++ {
            if dp[j] != -1 && abs(nums[j]-nums[i]) <= target {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }

	if dp[n-1] == -1 {
		return -1
	}

	return dp[n-1]
}
