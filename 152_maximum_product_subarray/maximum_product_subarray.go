package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var ans int = nums[0]
	// 以i结尾的子数组的最大积
	maxDP := make([]int, len(nums))
	// 以i结尾的子数组的最小积
	minDP := make([]int, len(nums))
	//初始化DP；
	maxDP[0] = nums[0]
	minDP[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minDP[i] = min(nums[i], min(nums[i]*minDP[i-1], nums[i]*maxDP[i-1]))
		maxDP[i] = max(nums[i], max(nums[i]*minDP[i-1], nums[i]*maxDP[i-1]))
		ans = max(ans, maxDP[i])
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

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	maxProduct([]int{-2, 3, -4})
}
