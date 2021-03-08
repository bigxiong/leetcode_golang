package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil && root.Right == nil {
			return root.Right
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		p := root.Right
		for p.Left != nil {
			p = p.Left
		}
        root.Val = p.Val
        root.Right = deleteNode(root.Right, p.Val)
	}
	return root
}

func main() {
	// [5,3,6,2,4,null,7]
}
