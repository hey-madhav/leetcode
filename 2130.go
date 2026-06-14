package main

func pairSum(head *ListNode) int {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	maxSum := 0
	for i := 0; i < len(values)/2; i++ {
		maxSum = max(maxSum, values[i]+values[len(values)-1-i])
	}

	return maxSum
}
