package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp []int = make([]int, len(nums))
	dp[0] = nums[0]
	var ans int = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
