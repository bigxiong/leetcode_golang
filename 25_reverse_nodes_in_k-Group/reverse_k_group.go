package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 以head开头， 翻转k个组链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var first *ListNode = head
	var tail *ListNode = head
	for i := 0; i < k; i++ {
		if tail == nil {
			return first
		}
		tail = tail.Next
	}

	newHead := reverseList(first, tail)
	// 尾结点连接下一部分
	first.Next = reverseKGroup(tail, k)
	return newHead
}

// 翻转first -> tail
func reverseList(first, tail *ListNode) *ListNode  {
	var prev, cur *ListNode = nil, first
	for cur != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	
}
