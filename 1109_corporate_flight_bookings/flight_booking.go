package main

import "fmt"

func corpFlightBookings(bookings [][]int, n int) []int {
	//var ans []int = make([]int, n)
	//for i := 0 ; i < len(bookings); i++{
	//	book := bookings[i]
	//	start := book[0]
	//	end := book[1]
	//	for j := start; j <= end; j++{
	//		ans[j-1] += book[2]
	//	}
	//}
	var ans []int = make([]int, n)
	diff := NewDifferenceArray(ans)
	for i := 0 ; i < len(bookings); i++{
		book := bookings[i]
		start := book[0]
		end := book[1]
		val := book[2]
		diff.Increment(start-1, end-1, val)
	}

	return diff.GetResult()
}

type DifferenceArray struct {
	diff []int
}

func NewDifferenceArray(nums []int) *DifferenceArray {
	var diff []int = make([]int, len(nums))

	diff[0] = nums[0]
	for i := 1 ; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return &DifferenceArray{diff:diff}
}

/* 给闭区间 [i,j] 增加 val（可以是负数）*/
func (d *DifferenceArray) Increment(i, j, val int) {
	d.diff[i] += val
	if j + 1 < len(d.diff) {
		d.diff[j + 1] -= val
	}
}

func (d *DifferenceArray) GetResult() []int {
	res := make([]int, len(d.diff))
	// 根据差分数组构造结果数组
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i - 1] + d.diff[i]
	}
	return res
}

func main()  {
	var data []int = []int{0,1,2,3,4,5}
	diff := NewDifferenceArray(data)
	diff.Increment(0,5,1)
	fmt.Println(diff.GetResult())
}
