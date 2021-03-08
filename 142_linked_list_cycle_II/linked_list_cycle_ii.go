package main

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visisted := make(map[*ListNode]bool)
	p := head
	for p != nil {
		if _, ok := visisted[p]; ok {
			return p
		}
		visisted[p] = true
		p = p.Next
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}
	// 上面的代码类似 hasCycle 函数
	if fast == nil || fast.Next == nil {
		// fast 遇到空指针说明没有环
		return nil
	}
	p := head
	for p != slow {
		p = p.Next
		fast = fast.Next
	}
	return p
}