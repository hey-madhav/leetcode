package main

func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})

	for i := range nums {
		set[nums[i]] = struct{}{}
	}

	longest := 1
	for _, num := range nums {
		if _, ok := set[num-1]; !ok {
			start := num
			length := 1

			for ok := true; ok; _, ok = set[start+1] {
				start++
				length++
			}

			longest = max(longest, length)
		}
	}

	return longest
}