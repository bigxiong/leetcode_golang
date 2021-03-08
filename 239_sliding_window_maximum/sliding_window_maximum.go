package main

import "fmt"

// method 1 : brute force
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	for i := k - 1; i < len(nums); i++{
		wStart := i - k + 1
		wEnd := i
		var max int = nums[wStart]
		for j := wStart+1 ; j <= wEnd; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		res = append(res, max)
	}
	return res
}

// method 2 : window使用BST结构


// method 3 : monotonicQueue

type MonotonicQueue []int

func (m *MonotonicQueue) Push(x int)  {
	for i := len(*m) - 1; i >= 0; i-- {
		if (*m)[i] < x {
			*m = (*m)[:i]
		}
	}
	*m = append(*m, x)
}

func (m *MonotonicQueue) Pop() int {
	top := (*m)[0]
	*m = (*m)[1:]
	return top
}

func (m MonotonicQueue) Max() int {
	return m[0]
}

func maxSlidingWindow2(nums []int, k int) []int {
	var res []int
	var queue MonotonicQueue
	for i := 0; i < len(nums); i++{
		if i < k-1 {
			(&queue).Push(nums[i])
		} else {
			// i = k-1
			(&queue).Push(nums[i])
			max := queue.Max()
			res = append(res, max)
			if max == nums[i-k+1] {
				(&queue).Pop()
			}
		}
	}
	return res
}


func main()  {
	data := []int{1,3,-1,-3,5,3,6,7}
	fmt.Println(maxSlidingWindow(data, 3))
}