package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var best int = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		x1 := nums[i]
		for left, right := i+1, len(nums)-1; left < right; {
			sum := x1 + nums[left] + nums[right]
			if abs(target-sum) < abs(target-best) {
				best = sum
			}
			if sum < target {
				left++
			} else if target < sum {
				right--
			} else {
				return sum
			}
		}
	}

	return best
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -1 * x
	}
}

func main() {

}
