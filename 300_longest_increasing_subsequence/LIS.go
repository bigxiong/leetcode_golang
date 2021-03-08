package main

func lengthOfLIS(nums []int) int {
	var dp []int = make([]int, len(nums))
	for i := 0; i < len(dp); i++{
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++{
		for j := 0; j < i; j++{
			if nums[j] < nums[i] {
				dp[i] =  max(dp[i], dp[j]+1)
			}
		}
	}

	var ans int
	for i := 0; i < len(dp); i++{
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
