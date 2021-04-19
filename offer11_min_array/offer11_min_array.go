package main

func minArray(nums []int) int {
	left := 0
	right := len(nums) - 1
	//if nums[right] >= nums[left] {
	//	return nums[left]
	//}
	/*
		输入：[3,4,5,1,2]
		输出：1
	*/
	for left <= right {
		mid := left + (right-left)/2
		if mid == 0 && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	minArray([]int{1, 3, 5})
}
