package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// bad code
	//dummyHead := &ListNode{Val:-1,Next:head}
	//cur := head
	//prev := dummyHead
	//for cur != nil {
	//	if cur.Val == Val {
	//		prev.Next = cur.Next
	//		cur.Next = nil
	//	}
	//	prev = prev.Next
	//	cur = prev.Next
	//}
	//return dummyHead.Next

	dummyHead := &ListNode{Val:-1,Next:head}
	cur := head
	prev := dummyHead
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return dummyHead.Next
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy
	for n > 0 {
		fast = fast.Next
		n--
	}

	var prev *ListNode
	//i := 0
	for fast != nil {
		//if i >= n {
		//	prev = slow
		//	slow = slow.Next
		//}
		prev = slow
		slow = slow.Next
		fast = fast.Next
	//	i++
	}
	prev.Next = prev.Next.Next
	slow.Next = nil
	return dummy.Next
}

func main()  {
	root := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{
							Val:  6,
							Next: &ListNode{
								Val:  7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	removeNthFromEnd(&root, 2)
}

func reverseString(s []byte)  {
	left := 0
	right := len(s)-1
	for ;left < right;  {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}