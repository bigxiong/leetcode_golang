package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	var ans = make([]int, int(math.Pow10(n)-1))
	for i := 0; i < len(ans); i++ {
		ans[i] = i + 1
	}
	return ans
}

func printNumbers2(n int) []int {
	var ans = make([]int, int(math.Pow10(n)-1))
	for digit := 1; digit < n+1; digit++ {

		// 第一位数字为1-9, 不能为0
		for first := '1'; first <= '9'; first++ {
			numsCharArray := make([]byte, digit)
			numsCharArray[0] = byte(first)
			dfs(1, numsCharArray, digit)
		}

	}
	return ans
}

// digit为数字位数
func dfs(index int, nums []byte, digit int) {
	// 填完了，开始打印
	if index == digit {
		fmt.Println(string(nums))
		return
	}
	for i := '0'; i <= '9'; i++ {
		nums[index] = byte(i)
		dfs(index+1, nums, digit)
	}
}

func main() {
	printNumbers2(4)
}
