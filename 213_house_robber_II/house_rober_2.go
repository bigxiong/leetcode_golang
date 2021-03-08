package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

// 偷[i...j]之间的房子
func robRange(nums []int, i, j int) int {
	// i-1, i-2
	dpi1, dpi2 := max(nums[i+1], nums[i]), nums[i]
	var dp int
	if j - i == 1 {
		return dpi1
	}
	for k := i+2; k <= j; k++ {
		dp = max(nums[k] + dpi2, dpi1)
		dpi2, dpi1 = dpi1, dp
	}
	return dp
}

func max(x,y int) int  {
	if x > y {
		return x
	} else {
		return y
	}
}

func main()  {
	nums := []int{2,7,9,3,1}
	fmt.Println(robRange(nums, 0, len(nums)-1))

	fmt.Println(rob([]int{2,3,2}))
	fmt.Println(rob([]int{0,0}))
}