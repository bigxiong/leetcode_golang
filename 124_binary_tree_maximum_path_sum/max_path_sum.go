package main

import "math"


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var ans int = math.MinInt32
	var onsideMax func(root *TreeNode) int
	onsideMax = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := onsideMax(root.Left)
		right := onsideMax(root.Right)
		ans = max(ans, left+right+root.Val)

		return max(left, right) + root.Val
	}

	onsideMax(root)
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	
}
