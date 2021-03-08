package main

import "math"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func maxDepth(p *TreeNode) int {
	if p == nil {
		return 0
	}
	var leftDepth, rightDepth int
	leftDepth = maxDepth(p.Left) +  1
	rightDepth = maxDepth(p.Right) + 1

	if leftDepth > rightDepth {
		return leftDepth
	} else {
		return rightDepth
	}
}

func isValidBST(root *TreeNode) bool {
	var f func(p *TreeNode, min, max int) bool
	f = func(p *TreeNode, min, max int) bool{
		if p == nil {
			return true
		}
		if p.Val <= min || p.Val >= max{
			return false
		}
		return f(p.Left, min, p.Val) && f(p.Right, p.Val, max)
	}

	return f(root, math.MinInt64, math.MaxInt64)
}

func main() {
	p := TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	maxDepth(&p)

}
