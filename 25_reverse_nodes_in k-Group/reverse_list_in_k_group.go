package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(a, b *ListNode) *ListNode {
	var prev *ListNode
	cur := a
	var next *ListNode
	for cur != b {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return a
		}
		b = b.Next
	}
	abHead := reverseList(a, b)
	a.Next = reverseKGroup(b, k)
	return abHead
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

