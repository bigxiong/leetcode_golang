package main

import "fmt"

func sortColors(nums []int)  {
	// 0, 1, 2
	buckets := make([]int, 3)

	for i := 0; i < len(nums); i++{
		// map index
		idx := nums[i]
		// counting
		buckets[idx]++
	}

	// sum count
	sum := 0
	for i, num := range buckets {
		sum += num
		buckets[i] = sum
	}

	sortedArray := make([]int, len(nums))
	for k := len(nums)-1; k >= 0; k--{
		val := nums[k]
		idx := buckets[val]
		fmt.Println(idx,":", val)
		sortedArray[idx-1] = val
		buckets[val]--
	}
	copy(nums, sortedArray)
}


func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}
