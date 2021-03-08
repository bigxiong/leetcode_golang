package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	var result int
	var backtrack func([]int, int, int)
	backtrack = func(nums []int, i, rest int) {
		if i == len(nums) {
			if rest == 0 {
				result++
			}
			return
		}

		// 添加一个+号
		rest -= nums[i]
		backtrack(nums, i+1, rest)
		rest += nums[i]


		rest += nums[i]
		backtrack(nums, i+1, rest)
		rest -= nums[i]
	}
	backtrack(nums, 0, S)
	return result
}

func findTargetSumWays2(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	var result int
	var backtrack func(nums []int, i int, curSum int, S int)
	backtrack = func(nums []int, i int, curSum int, S int) {
		if i == len(nums) {
			// 和为S
			if curSum == 0 {
				result++
			}
			return
		}
		// 使用加号
		backtrack(nums, i+1, curSum-nums[i], S)
		// 使用减号
		backtrack(nums, i+1, curSum+nums[i], S)

	}

	backtrack(nums, 0, 0, S)
	return result
}

/**
动态规划
+ sum(p)
- sum(n)

=> sum(p)-sum(n) = S
=> sum(p) = S + sum(n)
=> sum(p) + sum(p) = S + sum(n) + sum(p)
=> 2sum(p) = S + sum(nums)
=> sum(p) = (S + sum(nums)) / 2 (一个定值)
=> 在nums找出一个子序列，使得sum 为 (S + sum(nums)) / 2
*/
func findTargetSumWays3(nums []int, S int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 所有的和都<S 或者 sum + S 等于奇数， 说明不存在一个子集满足条件
	if sum < S || (sum + S) % 2 == 1 {
		return 0
	}

	return subsets(nums, (sum + S) / 2)
}

func subsets(nums []int, target int) int {
	r := len(nums)
	c := target + 1
	dp := make([][]int, r+1)
	for i := range dp {
		dp[i] = make([]int, c+1)
	}
	// 如果要凑成target=0, 那么什么都不装是唯一的方法
	for i := 0 ; i <= r; i++ {
		dp[i][0] = 1
	}
	// 如果 数字个数为0， 那么无法凑成target
	for j := 0; j <= c; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= r; i++ {
		for j := 1; j <= c; j++ {
			// 当前数字 < target， 可以装入背包
			if nums[i-1] <= j {
				// 可以选择装或者不装
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			// 不可装入，那么只有一种选择
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	// 前r个数，凑成sum为c的方法数
	return dp[r][c]
}

func main() {
	s := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWays2(s, 3))
}
