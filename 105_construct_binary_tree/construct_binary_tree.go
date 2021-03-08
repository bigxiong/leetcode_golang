package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func([]int, []int, int, int, int, int) *TreeNode
	index := make(map[int]int)
	build = func(preorder []int, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}
		rootVal := preorder[preStart]
		root := &TreeNode{rootVal, nil, nil}
		rootIndex := index[rootVal]
		//for ; rootIndex <= inEnd; rootIndex++ {
		//	if inorder[rootIndex] == rootVal {
		//		break
		//	}
		//}
		leftLen := rootIndex - inStart
		root.Left  = build(preorder, inorder, preStart+1, preStart+leftLen, inStart, rootIndex-1)
		root.Right = build(preorder, inorder, preStart+leftLen+1, preEnd, rootIndex+1, inEnd)
		return root
		return nil
	}


	for i, val := range inorder {
		index[val] = i
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func main() {
	
}
