package main

func combine(n int, k int) [][]int {
	var backtrack func(n int, nums []int, start, k int)
	var res [][]int
	backtrack = func(n int, nums []int, start int, k int) {
		if len(nums) == k {
			// 此处需要copy和return 结束递归
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			res = append(res, tmp)
			return
		}
		for i := start; i <= n; i++ {
			nums = append(nums, i)
			// 这里是i+1, 从下一个开始
			backtrack(n, nums, i+1, k)
			nums = nums[:len(nums)-1]
		}
	}
	backtrack(n, nil, 1, k)
	return res
}

func main() {
	combine(4, 2)
}
