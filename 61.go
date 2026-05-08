package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	length := 1
	tail := head

	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	k = k % length
	if k == 0 {
		return head
	}

	breakpoint := length - k

	current := head
	for i := 1; i < breakpoint; i++ {
		current = current.Next
	}

	newHead := current.Next
	current.Next = nil
	tail.Next = head

	return newHead
}
