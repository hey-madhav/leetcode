package main

func mirrorDistance(n int) int {
	return abs(n - reverse(n))
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func reverse(num int) int {
	var rev = 0
	for num > 0 {
		remainder := num % 10
		rev = 10*rev + remainder
		num = num / 10
	}

	return rev
}