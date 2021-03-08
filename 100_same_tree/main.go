package main


// Definition for a binary tree node

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// DFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)

}
// BFS
func isSameTree2(p *TreeNode, q *TreeNode) bool{
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) != 0 && len(queue2) != 0 {
		// 查看值是否相等
		node1 := queue1[0]
		node2 := queue2[0]
		if node1.Val != node2.Val {
			return false
		}

		queue1 = queue1[1:]
		queue2 = queue2[1:]

		left1, right1 := node1.Left, node1.Right
		left2, right2 := node2.Left, node2.Right
		if (left1 == nil && left2 != nil ) || (left1 != nil && left2 == nil ) {
			return false
		}
		if (right1 == nil && right2 != nil ) || (right1 != nil && right2 == nil ) {
			return false
		}
		if left1 != nil {
			queue1 = append(queue1, left1)
		}
		if right1 != nil {
			queue1 = append(queue1, right1)
		}
		if left2 != nil {
			queue2 = append(queue2, left2)
		}
		if right2 != nil {
			queue2 = append(queue2, right2)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}

func main() {
	
}
