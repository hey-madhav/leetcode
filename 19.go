package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var listSize int
	current := head
	for current != nil {
		listSize++
		current = current.Next
	}

	if n > listSize {
		return head
	}

	indexToRemove := listSize-n

	if indexToRemove == 0 {
		return head.Next
	}

	current = head
	for i:=0;i<indexToRemove-1;i++ {
		current = current.Next
	}

	current.Next = current.Next.Next

	return head
}