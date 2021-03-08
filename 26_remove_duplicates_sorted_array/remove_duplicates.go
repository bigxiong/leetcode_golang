package main

import "fmt"

func removeDuplicates(nums []int) int {
	var lastNotDuplicated int = 0
	for j := 1; j < len(nums); j++{
		// 这里不要写成for(多此一举), 考虑相反情况
		//for j < len(nums) && nums[j] == nums[lastNotDuplicated] {
		//	j++
		//}
		// 不相等，赋值。 否则继续j++往后走
		if nums[j] != nums[lastNotDuplicated] {
			nums[lastNotDuplicated+1] = nums[j]
			lastNotDuplicated++
		}
	}
	return lastNotDuplicated+1
}


func main() {
	nums := []int{1,1,1,1,2,3,3}
	fmt.Println(removeDuplicates(nums))
}
