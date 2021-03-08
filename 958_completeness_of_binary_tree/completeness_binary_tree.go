package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type NumberTreeNode struct {
	Number int
	Node *TreeNode
}
// BFS 通过层次遍历来编号
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*NumberTreeNode
	queue = append(queue, &NumberTreeNode{1, root})
	var i int = 0
	for i < len(queue) {
		node := queue[i]
		// 非空节点
		if node.Node != nil {
			// left child number is 2*i （may be nil）
			queue = append(queue, &NumberTreeNode{2*node.Number, node.Node.Left})
			// right child number is 2*i+1 （may be nil）
			queue = append(queue, &NumberTreeNode{2*node.Number+1, node.Node.Right})
		}
		i++
	}
	return queue[i-1].Number == len(queue)
}


func main() {
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	isCompleteTree(&root)
}
