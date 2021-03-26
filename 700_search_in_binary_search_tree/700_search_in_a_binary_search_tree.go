package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// if val < root.Val {
	//     return searchBST(root.Left, val)
	// } else if root.Val < val {
	//     return searchBST(root.Right, val)
	// } else {
	//     return root
	// }
	p := root
	for p != nil {
		if val < p.Val {
			p = p.Left
		} else if p.Val < val {
			p = p.Right
		} else {
			return p
		}
	}
	return nil
}

func main() {
	// 40,20,60,10,30,50,70
	root := TreeNode{
		Val: 40,
		Left: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   30,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 60,
			Left: &TreeNode{
				Val:   50,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   70,
				Left:  nil,
				Right: nil,
			},
		},
	}

	searchBST(&root, 25)
}
