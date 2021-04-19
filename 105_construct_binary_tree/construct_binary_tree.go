package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
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
		root.Left = build(preorder, inorder, preStart+1, preStart+leftLen, inStart, rootIndex-1)
		root.Right = build(preorder, inorder, preStart+leftLen+1, preEnd, rootIndex+1, inEnd)
		return root
	}

	for i, val := range inorder {
		index[val] = i
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	var build func([]int, []int) *TreeNode
	index := make(map[int]int)
	build = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {
			return nil
		}
		rootVal := preorder[0]
		root := &TreeNode{rootVal, nil, nil}

		rootIdx := 0
		for i := range inorder {
			if inorder[i] == rootVal {
				rootIdx = i
				break
			}
		}
		root.Left = build(preorder[1:rootIdx+1], inorder[:rootIdx])
		root.Right = build(preorder[rootIdx+1:], inorder[rootIdx+1:])
		return root
	}

	for i, val := range inorder {
		index[val] = i
	}
	return build(preorder, inorder)
}

func main() {

}
