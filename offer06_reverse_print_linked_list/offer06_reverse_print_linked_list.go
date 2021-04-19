package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var ans []int
	var doReversePrint func(head *ListNode)
	doReversePrint = func(head *ListNode) {
		if head == nil {
			return
		}
		doReversePrint(head.Next)
		ans = append(ans, head.Val)
	}

	doReversePrint(head)
	return ans
}

func main() {

}
