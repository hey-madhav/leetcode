package main

func minOperations(nums []int) int64 {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	temp := nums
	cost := int64(0)
	prevMax := temp[0]
	prevReq := 0

	for i := 1; i < n; i++ {
		if temp[i] > prevMax {
			prevMax = temp[i]
		}

		req := 0
		if prevMax > temp[i] {
			req = prevMax - temp[i]
		}

		if req > prevReq {
			cost += int64(req - prevReq)
		}
		prevReq = req
	}

	return cost
}
