package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode)  {
	// 1. get inorder numbers
	var inorder []int
	var inorderFunc func(root *TreeNode)
	inorderFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderFunc(root.Left)
		inorder = append(inorder, root.Val)
		inorderFunc(root.Right)
	}
	inorderFunc(root)

	// 2. find swapped
	/**
	[1,2,3,4,5,6,7]
	[1,6,3,4,5,2,7]
	*/
	x, y := -1, -1
	for i := 0; i < len(inorder); i++{
		if inorder[i] > inorder[i+1] {
			y = inorder[i+1]
			if x == -1 {
				x = inorder[i]
			} else {
				break
			}
		}
	}
	fmt.Println("swapped: ", x, y)
	recover(root, 2, x, y)
}

func recover(root *TreeNode, count int, x int, y int)  {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else if root.Val == y {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(root.Left, count, x, y)
	recover(root.Right, count, x, y)
}


func recoverTree2(root *TreeNode)  {
	var prev *TreeNode = &TreeNode{math.MinInt64, nil, nil}
	var x *TreeNode
	var y *TreeNode
	var f func(root *TreeNode)
	f = func(curl *TreeNode) {
		if curl == nil {
			return
		}
		f(curl.Left)
		if x == nil && prev.Val > curl.Val {
			x = prev
		}
		if x != nil && prev.Val > curl.Val {
			y = curl
		}
		prev = curl
		f(curl.Right)
	}
	f(root)

	// do swap
	x.Val, y.Val = y.Val, x.Val
}

func main() {
	
}
