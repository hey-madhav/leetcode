package main

import "math"

func minMirrorPairDistance(nums []int) int {
	var minimumDistance = math.MaxInt
	revIndex := make(map[int]int)

	for i, num := range nums {
        if prev, ok := revIndex[num]; ok {
            distance := i - prev
            minimumDistance = min(minimumDistance, distance)
        }
        revIndex[reverse(num)] = i
    }

	if minimumDistance == math.MaxInt {
		return -1
	}

	return minimumDistance
}

// func reverse(num int) int {
// 	var rev = 0
// 	for num > 0 {
// 		remainder := num % 10
// 		rev = 10*rev + remainder
// 		num = num / 10
// 	}

// 	return rev
// }
