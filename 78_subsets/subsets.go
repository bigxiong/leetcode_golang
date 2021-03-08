package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	var backtrack func(nums []int, track []int, k int)
	backtrack = func(nums []int, track []int, k int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ans = append(ans, tmp)
		ans = append(ans, track)
		for i := k ; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums, track, i+1)
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, nil, 0)
	return ans

}

func main()  {
	fmt.Println(subsets([]int{1,2,3}))
}