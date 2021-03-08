package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

var successor *ListNode
func reverseListN(head *ListNode, n int) *ListNode  {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}

func reverseListRecursive(head *ListNode) *ListNode  {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func reverseBetweenRecursive(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseListN(head, n)
	}
	head.Next = reverseBetweenRecursive(head.Next, m - 1, n - 1)
	return head
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead := ListNode{
		Val:  -1,
		Next: head,
	}
	prev := &dummyHead
	cur  := head
	// [1]. 走到第m个节点。
	for m > 1 {
		prev = cur
		cur = cur.Next
		n--
		m--
	}
	// [2]. 记录断开点信息
	con := prev // 第一部分最后一个节点
	tail := cur // 第二部分尾节点（反前转头节点）

	// [3]. 开始反转操作
	// 经典反转链表部分。
	for n > 0 {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		n--
	}
	// [4]. 组装连接链表
	// [反转后]
    // prev为反转后头节点(第二部分头节点)
    // cur为第三部分头节点
    // [连接]
    // 第一部分 --> 第二部分      --> 第三部分
    // ----con ->  prev---tail  --> cur
    con.Next = prev
	tail.Next = cur

	return dummyHead.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}


/**
[ 3 , 4 , 5  6 , 9 , 4 ]


func(*head, k int32)

eg: k = 2

4 3 4 9 6 5


eg: k = 3

5 4 3 4 9 6



*/
func reverseFromKthNode(head *ListNode, k int32) *ListNode{
	if k == 0 {
		return reverseList(head)
	}

	firtTail := head
	var prev *ListNode
	cur := head
	for k > 0  {
		if cur == nil {
			return reverseList(head)
		}
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
		k--
	}
	// 5, 4 , 3
	newHead := prev

	prev = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}


	firtTail.Next = prev
	return newHead

}

func main() {
	// 1->2->3->4->5->6->->7->NULL
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
						Next: nil,
					},
				},
			},
		},
	}
	//root = ListNode{
	//	Val:  3,
	//	Next: &ListNode{
	//		Val:  5,
	//		Next: nil,
	//	},
	//}
	// reverseBetween(&root, 1, 2)
	//reverseFromKthNode(&root, 3)
	reverseListN(&root, 3)
}
