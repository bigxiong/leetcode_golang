package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	var backtrack func(nums []int, track []int, start int)
	backtrack = func(nums []int, track []int, start int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ans = append(ans, tmp)
		// 此处添加条件len(track) == k，可以打印不同的k个数的组合。 leetcode 77

		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums, track, i+1)
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, nil, 0)
	return ans

}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	//subsets([]int{1,2,3})
}
