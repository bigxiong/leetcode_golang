package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return  nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	level := 1
	var results [][]int
	for len(queue) > 0 {
		size := len(queue)
		levelRes := make([]int, size)
		for  i := 0; i < size; i++ {
			// pop one element
			top := queue[0]
			queue = queue[1:]
			if level % 2 != 0 {
				levelRes[i] = top.Val
			} else {
				levelRes[size-i-1] = top.Val
			}
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		results = append(results, levelRes)
		level++
	}
	return results
}

func main() {
	
}
