package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	hashTable := make(map[int]struct{})
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	prev := dummy
	for prev.Next != nil {
		cur := prev.Next
		if _, ok := hashTable[cur.Val]; ok {
			prev.Next = cur.Next
			cur.Next = nil
			cur = prev.Next
		} else {
			hashTable[cur.Val] = struct{}{}
			prev = prev.Next

		}
	}
	return dummy.Next
}


func main() {
	// [1, 2, 3, 3, 2, 1]

	p := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
		},
	}

	removeDuplicateNodes(&p)

}
