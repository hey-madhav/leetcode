package main

import "sort"

type frequency struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
	var frequencyCount = make(map[int]int)

	for i := range nums {
		frequencyCount[nums[i]]++
	}

	var frequencyStructArray []frequency
	for key, value := range frequencyCount {
		frequencyStructArray = append(frequencyStructArray, frequency{key, value})
	}

	sort.Slice(frequencyStructArray, func(i, j int) bool {
		return frequencyStructArray[i].Value > frequencyStructArray[j].Value
	})

	var result []int
	for i:=0;i<k;i++ {
		result = append(result, frequencyStructArray[i].Key)
	}

	return result
}
