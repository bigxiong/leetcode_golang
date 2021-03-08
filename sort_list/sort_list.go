package main

import (
    "container/heap"
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    // 1. cur list O(n)
    middle := getListMiddle(head)
    if middle == head && head.Next == nil {
        return head
    }
    p1 := head
    p2 := middle.Next
    middle.Next = nil

    p1 = sortList(p1)
    p2 = sortList(p2)

    return mergeTwoSortedList(p1, p2)
}

func getListMiddle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    fast := head
    slow := head
    pre := slow
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    return pre
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

func printList(head *ListNode) {
    for head != nil {
        fmt.Printf("%d->",head.Val)
        head = head.Next
    }
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode{
    if left == right {
        return lists[left]
    }
    m := left + (right-left)/2
    p1 := merge(lists, left, m)
    p2 := merge(lists, m+1, right)
    return mergeTwoSortedList(p1, p2)
}

// 使用优先级队列
func mergeKLists2(lists []*ListNode) *ListNode {
    var pq  HeapMin
    heap.Init(&pq)
    // O(klogk)
    for _, list := range lists {
        if list != nil {
            heap.Push(&pq, list)
        }
    }

    dummy := ListNode{
        Val:  -1,
        Next: nil,
    }
    cur := &dummy

    for pq.Len() > 0 {
        min := heap.Pop(&pq).(*ListNode)
        cur.Next = min
        cur = cur.Next
        if min.Next != nil {
            heap.Push(&pq, min.Next)
        }
    }
    return dummy.Next
}

type HeapMin []*ListNode


// Len is the number of elements in the collection.
func (h HeapMin) Len() int {
    return len(h)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (h HeapMin) Less(i, j int) bool {
   return h[i].Val < h[j].Val
}
// Swap swaps the elements with indexes i and j.
func (h *HeapMin) Swap(i, j int) {
    (*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
// add x as element Len()
func (h *HeapMin) Push(x interface{})  {
    item := x.(*ListNode)
    *h = append(*h, item)
}
// remove and return element Len() - 1.
func (h *HeapMin) Pop() interface{} {
    old := *h
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    *h = old[0:n-1]
    return item
}

func main() {
	head := ListNode{
        Val:  4,
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
    }
    printList(&head)
    fmt.Println()

	sorted := sortList(&head)
    printList(sorted)
    // 1->4->3->2->5->NULL
}
