package main

func twoSum(nums []int, target int) []int {
	var indexMap = make(map[int]int)

	for i := range nums {
		if ind, ok := indexMap[target-nums[i]]; ok {
			return []int{ind, i}
		}

		indexMap[nums[i]] = i
	}

	return nil
}
