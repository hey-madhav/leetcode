package main

import "math"

func reverse(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}

func rotateLeft(arr []int) []int {
	start := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}

	arr[len(arr)-1] = start
	return arr
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}

	return true
}

func minOperations(nums []int) int {
	if isSorted(nums) {
		return 0
	}

	ans := math.MaxInt

	arr := append([]int{}, nums...)

	for i := 1; i <= len(nums); i++ {
		arr = rotateLeft(arr)
		if isSorted(arr) {
			ans = min(ans, i)
			break
		}
	}

	rev := reverse(append([]int{}, nums...))

	if isSorted(rev) {
		ans = min(ans, 1)
	}

	arr = append([]int{}, rev...)
	for i := 1; i <= len(nums); i++ {
		arr = rotateLeft(arr)
		if isSorted(arr) {
			ans = min(ans, i+1)
			break
		}
	}

	if ans != math.MaxInt {
		return ans
	}

	return -1
}
