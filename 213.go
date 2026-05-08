package main

func rob2(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    if n == 2 {
        if nums[0] > nums[1] {
            return nums[0]
        }
        return nums[1]
    }

    // Exclude last house
    nums1 := nums[:n-1]
    // Exclude first house
    nums2 := nums[1:]

    return max(robLinear(nums1), robLinear(nums2))
}

func robLinear(nums []int) int {
    prev1, prev2 := 0, 0
    for _, num := range nums {
        temp := prev1
        if prev2+num > prev1 {
            prev1 = prev2 + num
        }
        prev2 = temp
    }
    return prev1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
