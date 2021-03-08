package main

func firstMissingPositive(nums []int) int {
	hashMap := make(map[int]int)
	for _, num := range nums{
		hashMap[num] = 1
	}

	for i := 1; i <= len(nums); i++{
		if _, ok := hashMap[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}

func firstMissingPositive2(nums []int) int {
	for i := 0; i < len(nums); i++{
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	for i := 0; i < len(nums); i++{
		num := int(abs(nums[i]))
		index := num-1
		if index < len(nums) {
			nums[index] = -1 * abs(nums[index])
		}
	}
	for i := 0 ; i < len(nums); i++{
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func abs(x int) int  {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	firstMissingPositive2([]int{1,0})
}
