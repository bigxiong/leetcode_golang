package main


/*
	Input:  nums = [1,2,2,4]
	Output: [2,3]
    1...n
*/
func findErrorNums(nums []int) []int {
	var dup, missing int
	for i := 0 ; i < len(nums); i++{
		index := abs(nums[i])-1
		if nums[index] < 0 {
			dup = nums[i]
		} else {
			nums[index] = -1 * nums[index]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			missing = nums[i]
		}
	}

	return []int{dup, missing}
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x*-1
	}
}


func findErrorNums2(nums []int) []int {
	var dup, missing int
	numFreq := make(map[int]int)
	for i := 0 ; i < len(nums); i++{
		numFreq[nums[i]]++
	}

	for i := 1; i <= len(nums); i++{
		if freq, ok := numFreq[i]; ok {
			if freq == 2 {
				dup = i
			}
		} else {
			missing = i
		}
	}

	return []int{dup, missing}
}

func main()  {
	findErrorNums2([]int{2,1,2,4})
}