package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumValueSum(nums []int, s string) int64 {
	h := &MinHeap{}
	heap.Init(h)

	var sum int64
	tokens := 0

	for i := len(nums) - 1; i >= 0; i-- {
		heap.Push(h, nums[i])
		sum += int64(nums[i])

		if s[i] == '1' {
			tokens++
		}

		for h.Len() > tokens {
			sum -= int64(heap.Pop(h).(int))
		}
	}

	return sum
}
