package main

import "fmt"

func findRepeatNumber(nums []int) int {
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	for k, v := range hashMap {
		if v > 1 {
			return k
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	hashSet := make(map[int]struct{})
	var repeat int
	for _, num := range nums {
		if _, ok := hashSet[num]; ok {
			repeat = num
			break
		}
		hashSet[num] = struct{}{}
	}
	return repeat
}

func main() {
	ans := findRepeatNumber([]int{1, 2, 3, 2, 2, 3, 4, 6, 7, 9})
	fmt.Println(ans)
}
