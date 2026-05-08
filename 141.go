package main

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = head.Next
		fast = head.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}