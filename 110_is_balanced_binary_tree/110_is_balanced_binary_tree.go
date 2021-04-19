package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := getHeightAndBalanced(root)
	return isBalanced
}

func getHeightAndBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftH, isLeftB := getHeightAndBalanced(root.Left)
	rightH, isRightB := getHeightAndBalanced(root.Right)
	if !isLeftB || !isRightB {
		return max(leftH, rightH) + 1, false
	}

	if abs(leftH-rightH) > 1 {
		return max(leftH, rightH) + 1, false
	}

	return max(leftH, rightH) + 1, true
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {

}
