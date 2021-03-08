package main

/*
[2, 3, 1, 1, 4]

[3, 2, 1, 0, 4]

*/
func canJump(nums []int) bool {
	var rightmost int = 0
	for i := 0; i < len(nums)-1; i++ {
		rightmost = max(rightmost, i + nums[i])
		// 可以跳到最后一行
		if rightmost <= i {
			return false
		}
	}
	return rightmost >= len(nums)-1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	
}
