package main

func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
	fast, slow := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

    if fast == slow {
        head.Next = nil
        return head
    }

    if fast.Next != nil {
        slow = slow.Next
    }

	temp := head
	for temp.Next != slow {
		temp = temp.Next
	}

	temp.Next = slow.Next

	return head
}
