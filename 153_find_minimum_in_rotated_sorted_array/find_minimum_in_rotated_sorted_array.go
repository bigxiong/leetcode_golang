package main

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[right] >= nums[left] {
		return nums[left]
	}
	for left <= right {
		mid := left + (right-left)/2
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
