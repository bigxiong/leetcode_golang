package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	m := left + (right-left)/2
	p1 := merge(lists, left, m)
	p2 := merge(lists, m+1, right)
	return mergeTwoSortedList(p1, p2)
}

func mergeTwoSortedList(p1 *ListNode, p2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}

func main() {

}
