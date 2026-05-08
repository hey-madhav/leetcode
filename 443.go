package main

import (
	"strconv"
)

func compress(chars []byte) int {
	left, right, currentCount, s := 0, 0, 0, ""

	for right < len(chars) {
		if chars[right] == chars[left] {
			currentCount++
		} else {
			s += string(chars[left]) + strconv.Itoa(currentCount)
			left = right
			currentCount = 1
		}
		right++
	}

	s += string(chars[left]) + strconv.Itoa(currentCount)
	return len(s)
}
