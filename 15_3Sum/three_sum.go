package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		x1 := nums[i]
		if i > 0 && x1 == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			if x1+nums[left]+nums[right] == 0 {
				res = append(res, []int{x1, nums[left], nums[right]})
				left++
				for left < len(nums) && nums[left] == nums[left-1] {
					left++
				}
				right--
				for right >= 0 && nums[right] == nums[right+1] {
					right--
				}
			} else if x1+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
