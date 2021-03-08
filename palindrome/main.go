package main

import (
	"fmt"
	"strings"
)

// 9. Palindrome Number
func isPalindrome(x int) bool {
	// 10
	if x % 10 == 0 && x != 0 {
		return false
	}
	var revertedNumber int
	for revertedNumber < x {
		revertedNumber = revertedNumber*10 + x%10
		x = x/10
	}
    // 11 --> 11
	return revertedNumber == x || revertedNumber/10 == x
}


func isalphanumeric(c rune) bool {
	return ('a' <= c && c <= 'z' ) ||  ('0' <= c && c <= '9' )
}

// 125. Valid Palindrome
func isPalindrome1(s string) bool {
	sArray := ([]rune)(strings.ToLower(s))
	for i, j := 0, len(s)-1; i < j ; {
		// continue 或者 使用for
		if !isalphanumeric(sArray[i]) {
			i++
			continue
		}
		if !isalphanumeric(sArray[j]) {
			j--
			continue
		}
		if sArray[i] != sArray[j] {
			return false
		}
		i++
		j--
	}
	return true
}


// 234. Palindrome Linked List


type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome3(head *ListNode) bool {
	var left *ListNode = head
	var traverse func(head *ListNode) bool
	traverse = func(head *ListNode) bool{
		if head == nil {
			return true
		}
		res := traverse(head.Next)
		if !res {
			return res
		}
		if head.Val != left.Val {
			return false
		}
		left = left.Next
		return res
	}
	return traverse(head)
}

func isPalindrome3_1(head *ListNode) bool {
	half := halfOfList(head)
	p := reverseList(half)
	q := head
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}
/*
      s     f
1->2->3->4->5->nil
      s     f
1->2->3->4->nil

*/
func halfOfList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast != nil {
		return slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode  {
	var prev, cur *ListNode = nil, head
	var next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}



func main() {
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(121121))

}
