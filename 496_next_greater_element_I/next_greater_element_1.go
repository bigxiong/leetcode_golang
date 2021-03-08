package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res []int
	var stack []int
	resultMap := make(map[int]int)
	for i := 0 ; i < len(nums2); i++{
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			resultMap[top] = nums2[i]
		}
		stack = append(stack, nums2[i])
	}
	 for len(stack) > 0 {
		 top := stack[len(stack)-1]
		 stack = stack[:len(stack)-1]
		 resultMap[top] = -1
	 }

	 for i := 0 ; i < len(nums1); i++{
	 	res = append(res, resultMap[nums1[i]])
	 }
	 return res
}


func nextGreaterElement2(nums []int) []int {
	var res []int = make([]int, len(nums))
	// 栈中保存的是右侧高度递增榜单
	var stack []int
	for i := len(nums)-1 ; i >= 0 ; i--{
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1]
		} else {
			res[i] = -1
		}

		stack = append(stack, nums[i])
	}

	return res
}

func main()  {
	d := []int{2,1,2,4,3}
	fmt.Println(nextGreaterElement2(d))
}