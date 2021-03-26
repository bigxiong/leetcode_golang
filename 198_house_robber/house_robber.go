package main

import "fmt"

// 1. brute force
func rob(nums []int) int {
	var robRecursive func(nums []int, start int) int
	// 偷到第k间房子，能获取的最大价值
	robRecursive = func(nums []int, k int) int {
		if k < 0 {
			return 0
		}
		// 1. 偷k, 则k-1不能偷，只能偷k-2
		x1 := nums[k] + robRecursive(nums, k-2)
		// 2. 不偷k, 则k-1可以偷
		x2 := robRecursive(nums, k-1)
		// 取2者较大值
		return max(x1, x2)
	}
	return robRecursive(nums, len(nums)-1)
}

// 使用memo数组，消除重复子问题
func rob1(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var robRecursive func(nums []int, start int) int
	// 偷到第k间房子，能获取的最大价值
	robRecursive = func(nums []int, k int) int {
		if k < 0 {
			return 0
		}
		if memo[k] != -1 {
			return memo[k]
		}
		// 1. 偷k, 则k-1不能偷，只能偷k-2
		x1 := nums[k] + robRecursive(nums, k-2)
		// 2. 不偷k, 则k-1可以偷
		x2 := robRecursive(nums, k-1)
		ans := max(x1, x2)
		memo[k] = ans
		return ans
	}
	return robRecursive(nums, 0)
}

// 动态规划（动态递归）
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))

	nums = []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
