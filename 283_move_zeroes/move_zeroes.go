package main

/**
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

func moveZeroes1(nums []int)  {
	var helperArray []int
	var zeroCount int
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			helperArray = append(helperArray, num)
		}
	}

	for i := 0; i < zeroCount; i++{
		helperArray = append(helperArray, 0)
	}

	copy(nums, helperArray)
}

func moveZeroes2(nums []int)  {
	var lastNotZero int
	for i := 0; i < len(nums); i++{
		if nums[i] != 0 {
			nums[lastNotZero] = nums[i]
			lastNotZero++
		}
	}
	for i := lastNotZero; i < len(nums); i++ {
		nums[i] = 0
	}
}


func moveZeroes3(nums []int)  {
	var lastNotZero int
	for i := 0; i < len(nums); i++{
		if nums[i] != 0 {
			nums[lastNotZero], nums[i] = nums[i], nums[lastNotZero]
			lastNotZero++
		}
	}
}

func main() {
	nums := []int{0,1,0,3,12}

	moveZeroes1(nums)

	nums = []int{0,1,0,3,12}
	moveZeroes2(nums)

	nums = []int{0,1,0,3,12}
	moveZeroes3(nums)
}
