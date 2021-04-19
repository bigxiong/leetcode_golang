package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{
		Val:  -1,
		Next: nil,
	}
	p := &dummyHead
	var tmp int = 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{
			Val:  tmp % 10,
			Next: nil,
		}
		p = p.Next
		tmp = tmp / 10
	}

	return dummyHead.Next
}
