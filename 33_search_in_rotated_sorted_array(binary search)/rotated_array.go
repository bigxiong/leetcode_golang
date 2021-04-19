package main

/*
                                   	 0 1 2 3 4 5 6 7 8
	[0,1,2,4,5,6,7,8] might become  [4,5,6,7,8,0,1,2,3])

*/
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[left] <= nums[right] {
		return nums[left]
	}
	for left < right {
		mid := (left + right) / 2
		// 错误的
		if mid > 0 && mid < len(nums)-1 && nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		if nums[left] <= nums[mid] && nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return nums[left]
}

func findMin2(nums []int) int {
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

func search(nums []int, target int) bool {
	if nums == nil {
		return false
	}
	left := 0
	end := len(nums) - 1
	var mid int
	for left <= end {
		mid = left + (end-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}
		if nums[left] < nums[mid] {
			//target在前半部分
			if nums[left] <= target && target < nums[mid] {
				end = mid - 1
			} else { //否则，去后半部分找
				left = mid + 1
			}
		} else {
			//后半部分有序
			//target在后半部分
			if nums[mid] < target && target <= nums[end] {
				left = mid + 1
			} else { //否则，去后半部分找
				end = mid - 1
			}
		}
	}
	return false
}

func search2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else {
			// 前半部分
			if nums[mid] >= nums[0] {
				if nums[left] <= target && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
				// 后半部分
			} else {
				if nums[mid] < target && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		}
	}
	return -1
}

func main() {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	findMin(a)
	//search2(a, 0)
	findMin2(a)
}
