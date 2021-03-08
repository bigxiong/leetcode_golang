package main

import "fmt"

/**
    1
   / \
  2   2
 / \ / \
3  4 4  3

*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root.Left, root.Right)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

func main() {
	// [1,2,2,null,3,null,3]
	p := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	r := isSymmetric(&p)
	fmt.Println(r)
}
