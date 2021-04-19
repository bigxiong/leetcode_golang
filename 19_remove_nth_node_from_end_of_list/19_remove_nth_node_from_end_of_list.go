package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	dummyHead.Next = head
	fast, slow := dummyHead, dummyHead
	for n >= 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func main() {

}
