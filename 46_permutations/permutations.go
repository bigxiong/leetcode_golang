package main

import "fmt"

func permute(nums []int) [][]int {
	var permutations [][]int
	var used map[int]bool = make(map[int]bool)
	var f func(nums []int, tracks []int)

	f = func(nums []int, tracks []int) {
		if len(tracks) == len(nums) {
			var dst []int = make([]int, len(tracks))
			copy(dst, tracks)
			fmt.Println(dst)
			// need copy
			permutations = append(permutations, dst)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			tracks = append(tracks, nums[i])
			used[nums[i]] = true
			f(nums, tracks)
			tracks = tracks[:len(tracks)-1]
			delete(used, nums[i])
		}
	}

	f(nums, nil)
	return permutations
}

func permute2(nums []int) [][]int {
	var permutations [][]int
	var f func(nums []int, first int)

	f = func(nums []int, first int) {
		if first == len(nums) {
			var dst []int = make([]int, len(nums))
			copy(dst, nums)
			fmt.Println(nums)
			// need copy
			permutations = append(permutations, dst)
			return
		}
		for i := first; i < len(nums); i++ {
			// swap
			nums[i], nums[first] = nums[first], nums[i]
			f(nums, first+1)
			// swap
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	f(nums, 0)
	return permutations
}

func main() {

	fmt.Println(permute2([]int{1, 2, 3}))
}
