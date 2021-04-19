package main

func findMin(nums []int) int {
	if nums == nil {
		return -1
	}

	left := 0
	right := len(nums) - 1
	// 处理全是升序的情形
	if nums[right] >= nums[left] {
		return nums[left]
	}
	/*
		输入：[3,4,5,1,2]
		输出：1
	*/
	for left <= right {
		mid := left + (right-left)/2
		// [3,1,2]
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[mid-1] {
			return nums[mid]
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

}
