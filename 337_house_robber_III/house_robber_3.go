package _37_house_robber_III


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rob1(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	var dp func(root *TreeNode) int
	// 前序遍历
	dp = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if val, ok := memo[root]; ok {
			return val
		}
		rob := root.Val
		if root.Left != nil {
			rob += dp(root.Left.Left) + dp(root.Left.Right)
		}
		if root.Right != nil {
			rob += dp(root.Right.Left) + dp(root.Right.Right)
		}

		notRob := dp(root.Left) + dp(root.Right)
		ans := max(rob, notRob)
		memo[root] = ans
		return ans
	}
	return dp(root)
}

func rob2(root *TreeNode) int {
	//memo := make(map[*TreeNode]int)
	// [0] 假如抢root 获取的最大值
	// [1] 假如不抢root 获取的最大值
	var dp func(root *TreeNode) []int
	// 前序遍历
	dp = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		var left []int = dp(root.Left)
		var right []int = dp(root.Right)

		// 1. rob root
		rob := root.Val + left[1] + right[1]
		// 2. not rob root
		notRob := max(left[0], left[1]) + max(right[0], right[1])

		return []int{rob, notRob}
	}
	ans := dp(root)
	return max(ans[0], ans[1])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}