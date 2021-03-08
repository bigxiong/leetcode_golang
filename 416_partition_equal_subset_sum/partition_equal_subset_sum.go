package main

import "fmt"

/**
转换为0-1 背包问题
=> 选取的数字的和恰好等于整个数组的元素和的一半
*/
func canPartition(nums []int) bool {
	var sum int = 0
	for _, num := range nums{
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	sum = sum / 2
	R := len(nums)
	C := sum
	dp := make([][]bool, R+1)
	for i := range dp {
		dp[i] = make([]bool, C+1)
	}
	// 背包空间为0时，相当于装满了
	for i := 0; i <= R ;i++ {
		dp[i][0] = true
	}

	for i := 1; i <= R; i++ {
		for curSum := 1; curSum <= C; curSum++ {
			// 1. 不能选择nums当前数字, 等于前i-1个数字，装满curSum 的值
			if (curSum - nums[i-1]) < 0 {
				dp[i][curSum] = dp[i-1][curSum]
			} else {
				// 可以选择加入或者不加入
				dp[i][curSum] = dp[i-1][curSum] || dp[i-1][curSum-nums[i-1]]
			}
			fmt.Printf("i=%v, sum=%v, dp[i][curSum]=%v \n", i, curSum, dp[i][curSum])
		}
	}
	return dp[R][C]
}

// 递归解法
func canPartition2(nums []int) bool {
	var sum int = 0
	for _, num := range nums{
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}
	sum = sum / 2

	var dp func(remainSum int, index int, nums []int) bool

	dp = func(remainSum int, N int, nums []int) bool {
		// base case
		if remainSum == 0 {
			return true
		}
		if remainSum < 0 {
			return false
		}
		if remainSum > 0 && N < 0 {
			return false
		}

		if remainSum - nums[N] < 0 {
			return dp(remainSum, N-1, nums)
		} else {
			return dp(remainSum, N-1, nums) || dp(remainSum-nums[N], N-1, nums)
		}
	}

	return dp(sum, len(nums)-1, nums)
}



func main() {
	fmt.Println(canPartition([]int{1,2,6,7,9,2,4,2,1}))
	result := canPartition2([]int{1,2,6,7,9,2,4,2,1})
	fmt.Println(result)
}
