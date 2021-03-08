package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var serialize func(root *TreeNode) string
	var hashMap = make(map[string]int)
	var solution []*TreeNode
	serialize = func(root *TreeNode) string {
		if root == nil {
			return "nil"
		}
		leftId := serialize(root.Left)
		rightId := serialize(root.Right)
		id := fmt.Sprintf("%d-%s-%s", root.Val, leftId, rightId)
		if times, ok := hashMap[id]; ok {
			if times == 1 {
				solution = append(solution, root)
				hashMap[id]++
			}
		} else {
			hashMap[id] = 1
		}
		return id
	}
	serialize(root)
	return solution
}

func main() {
	
}
