package main

import (
	"fmt"
	"math"
)

// 1. 递归求解(遍历所有情况)
func jump1(nums []int) int {
	return dp1(nums, 0)
}
// 从i跳到最后, 需要的最少步数
func dp1(nums []int, start int) int  {
	if start >= len(nums) - 1 {
		return 0
	}
	// 可选择的跳跃步数为(1,2,...nums[i])
	var res int = math.MaxInt32
	for i := 1; i <= nums[start]; i++{
		//选择跳跃i步
		subProblem := dp1(nums, start+i)
		res = min(res, subProblem+1)
	}
	return res
}


// 2. 记忆化递归求解
func jump2(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = math.MaxInt32
	}
	var dp func([]int, int) int
	dp = func(nums []int, start int) int  {
		if start >= len(nums) - 1 {
			return 0
		}
		if memo[start] != math.MaxInt32 {
			return memo[start]
		}

		// 可选择的跳跃步数为(1,2,...nums[i])
		for i := 1; i <= nums[start]; i++{
			//选择跳跃i步
			subProblem := dp(nums, start+i)
			memo[start] = min(memo[start], subProblem+1)
		}
		return memo[start]
	}

	return dp(nums, 0)
}


// 3. Greedy
// [2, 3, 1, 1, 4]
// i=0, end=2
// i=3, end=4
func jump3(nums []int) int {
	var rightmost int = 0
	var end int = 0 //最少跳跃次数， 当前能跳跃到的最远距离！ (i到end之间可以一步到达)
	var res int = 0 //跳跃次数
	for i := 0; i < len(nums)-1; i++{
		// 当前能跳到的最远距离
		rightmost = max(rightmost, i+nums[i])
		if end == i {
			res += 1        // 跳跃一次， 到达最远距离至rightmost
			end = rightmost
			if end >= len(nums)-1 {
				return res
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	res1 := dp1(nums, 3)
	res2 := dp1(nums, 2)
	res3 := dp1(nums, 1)
	res4 := dp1(nums, 0)
	fmt.Println(res1, " ", res2, " ", res3, " ", res4)
	fmt.Println(jump2(nums))
}
