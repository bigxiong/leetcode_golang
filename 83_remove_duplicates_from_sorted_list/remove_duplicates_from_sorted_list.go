package main


type ListNode struct {
     Val int
     Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode
	for fast != nil {
		if slow.Val != fast.Val {
			// 1. 直接赋值法
			//slow = slow.Next
			//slow.Val = fast.Val

			// 2. 链表指针操作
			// 清除野指针
			if slow.Next.Val != fast.Val {
				prev.Next = nil
			}
			slow.Next = fast
			slow = slow.Next
		}
		prev = fast
		fast = fast.Next
	}
	slow.Next = nil
	return head
}


func main()  {
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	deleteDuplicates(&head)
}
