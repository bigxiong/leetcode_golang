package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var ans [][]int
	for len(queue) > 0 {
		size := len(queue)
		var level []int
		for i := 0; i < size; i++ {
			front := queue[0]
			queue = queue[1:]
			level = append(level, front.Val)
			if front.Left != nil {
				queue = append(queue, front.Left)
			}
			if front.Right != nil {
				queue = append(queue, front.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}

func main() {
	// [3,9,20,null,null,15,7]
	root := TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	root.Left = &TreeNode{
		Val:   9,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val:   20,
		Left:   &TreeNode{
			Val:   15,
		},
		Right: &TreeNode{
			Val:   7,
		},
	}
	levelOrder(&root)
}
