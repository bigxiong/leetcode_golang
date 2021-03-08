package main

import (
	"container/list"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InsertIntoBST(root *TreeNode, data int) *TreeNode {
	node := TreeNode{
		Val:   data,
		Left:  nil,
		Right: nil,
	}
	p := root
	if p == nil {
		return &node
	}
	for {
		if data <= p.Val {
			if p.Left == nil {
				p.Left = &node
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &node
				break
			}
			p = p.Right
		}
	}
	return p

}

func InsertIntoBST2(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   data,
			Left:  nil,
			Right: nil,
		}
	}

	if data <= root.Val {
		root.Left = InsertIntoBST2(root.Left, data)
	} else {
		root.Right = InsertIntoBST2(root.Right, data)
	}
	return root
}

func preOrderTraversal(root *TreeNode) []int {
	stack := list.New()
	p := root
	var res []int
	for p != nil || stack.Len() != 0 {
		for p != nil {
			res = append(res, p.Val)
			stack.PushBack(p)
			p = p.Left
		}
		if stack.Len() != 0 {
			top := stack.Back()
			p = top.Value.(*TreeNode).Right
			stack.Remove(top)
		}
	}
	return res
}

func inOrderTraversal(root *TreeNode) []int {
	stack := list.New()
	p := root
	var res []int
	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		if stack.Len() != 0 {
			top := stack.Back()
			stack.Remove(top)
			e := top.Value.(*TreeNode)
			res = append(res, e.Val) // visit
			p = e.Right
		}
	}
	return res
}

var lastVisited *TreeNode
func postOrderTraversal(root *TreeNode) []int {
	stack := list.New()
	p := root
	var res []int
	for p != nil || stack.Len() != 0 {
		for p != nil {
			stack.PushBack(p)
			p = p.Left
		}
		if stack.Len() != 0 {
			top := stack.Back()
			e := top.Value.(*TreeNode)
			if (e.Left == nil && e.Right == nil) ||(e.Left != nil && e.Right == nil)|| e.Right == lastVisited {
				res = append(res, e.Val) // visit
				lastVisited = e
				stack.Remove(top)
			} else {
				p = e.Right
			}
		}
	}
	return res
}

func levelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var level []int
		for i := 0 ; i < size; i++ {
			f := queue[i] // pop
			level = append(level, f.Val)
			if f.Left != nil {
				queue = append(queue, f.Left)
			}
			if f.Right != nil {
				queue = append(queue, f.Right)
			}
		}
		ret = append(ret, level)
		// next level
		queue = queue[size:]
	}
	return ret
}

// 判断是否为平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isBalanced := getHeightAndBalanced(root)
	return isBalanced
}

func getHeightAndBalanced(p *TreeNode) (int, bool){
	if p == nil {
		return 0, true
	}
	leftHeight, leftBalanced := getHeightAndBalanced(p.Left)
	rightHeight,rightBalanced := getHeightAndBalanced(p.Right)
	if !leftBalanced || !rightBalanced {
		return 0, false
	}
	if abs(leftHeight-rightHeight) > 1 {
		return max(leftHeight, rightHeight)+1, false
	}

	return max(leftHeight, rightHeight)+1, true
}

func max(x,y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return x*-1
	}
}



func main() {
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	postOrderTraversal(&root)
}
