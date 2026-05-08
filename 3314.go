package main

func minBitwiseArray(nums []int) []int {
	var ans = make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}

	for i, num := range nums {
		for j := 1; j < num; j++ {
			if j | (j+1) == num {
				ans[i] = j
				break
			}
		}
	}

	return ans
}
