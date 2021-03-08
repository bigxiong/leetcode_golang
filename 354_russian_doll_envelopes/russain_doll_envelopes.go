package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	var dp []int
	var piles int

	for i := 0; i < len(nums); i++{
		//index := sort.SearchInts(dp, nums[i])
		index := BinarySearchLT(dp, nums[i])
		if index == -1 {
			dp = append(dp, nums[i])
			piles++
		} else {
			dp[index+1] = nums[i]
		}

	}
	return piles
}

type TwoDimensionArray [][]int

func (p TwoDimensionArray) Len() int { return len(p) }

func (p TwoDimensionArray) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] > p[j][1]
	} else {
		return p[i][0] < p[j][0]
	}
}

func (p TwoDimensionArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func maxEnvelopes(envelopes [][]int) int {
	sort.Sort(TwoDimensionArray(envelopes))
	fmt.Println(envelopes)
	var height []int = make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++{
		height[i] = envelopes[i][1]
	}
	return lengthOfLIS(height)
}


func BinarySearchLT(a []int, v int) int {
	left := 0
	right := len(a)-1
	for left <= right {
		mid := left + (right - left)/2
		if v < a[mid] {
			right = mid - 1
		} else if a[mid] < v {
			left = mid + 1
		} else {
			if mid != 0 && a[mid-1] < v{
				return mid-1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	array := [][]int{
		{5,4},
		{6,4},
		{6,7},
		{2,3},
	}
	fmt.Println(maxEnvelopes(array))
}
