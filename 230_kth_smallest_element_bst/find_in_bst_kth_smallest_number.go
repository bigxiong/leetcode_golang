package main

type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}


func kthSmallest(root *TreeNode, k int) int {
	var traverse func(root *TreeNode)
	var rank int
	var res int
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		rank++
		if rank == k {
			res = root.Val
			return
		}
		traverse(root.Right)
	}

	traverse(root)
	return  res
}

func main() {
	
}
