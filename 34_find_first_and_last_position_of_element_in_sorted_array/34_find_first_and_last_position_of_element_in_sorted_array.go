package main

func searchRange(nums []int, target int) []int {
	x1 := findFirst(nums, target)
	x2 := findLast(nums, target)
	return []int{x1, x2}
}

func findFirst(nums []int, x int) int {
	if nums == nil {
		return -1
	}
	left, right := 0, len(nums)-1
	// 1.此处是等于号
	for left <= right {
		mid := left + (right-left)/2
		//2. left, right  每次变换到mid位置
		if nums[mid] == x {
			// 2. 等于0的话，直接return
			if mid == 0 || nums[mid-1] != nums[mid] {
				return mid
			}
			right = mid - 1
		} else if x < nums[mid] {
			right = mid - 1
		} else if nums[mid] < x {
			left = mid + 1
		}
	}

	return -1
}

func findLast(nums []int, x int) int {
	if nums == nil {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == x {
			if mid == len(nums)-1 || nums[mid+1] != nums[mid] {
				return mid
			}
			left = mid + 1
		} else if x < nums[mid] {
			right = mid - 1
		} else if nums[mid] < x {
			left = mid + 1
		}
	}

	return -1
}

func main() {

}
