package main

import "fmt"
// method 1 : brute force
// O(n^3)
func subarraySum1(nums []int, k int) int  {
	var ans int
	for left := 0 ; left < len(nums); left++ {
		for right := left; right < len(nums); right++ {
			var sum int
			for i := left; i <= right; i++{
				sum += nums[i]
			}

			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// method 2 : brute force
// O(n^3)
func subarraySum2(nums []int, k int) int  {
	var ans int
	for left := 0 ; left < len(nums); left++ {
		var sum int
		for right := left; right < len(nums); right++ {
			//for i := left; i <= right; i++{
			//	sum += nums[i]
			//}
			sum += nums[right]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}


// method 2 : prefix sum
// O(n^3)
func subarraySum3(nums []int, k int) int {
	var prefixSum []int = make([]int,  len(nums)+1)
	var count int
	prefixSum[0] = 0
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	for left := 0 ; left < len(nums); left++ {
		for right := left ; right < len(nums); right++ {
			// [left, right] 子数组
			if prefixSum[right+1] - prefixSum[left] == k {
				count++
			}
		}
	}

	return count
}


// method 4: hash map 优化

func subarraySum(nums []int, k int) int {
	var result int
	preSumFreq := make(map[int]int)
	preSumFreq[0] = 1
	preSum := 0
	for i := 0 ; i < len(nums); i++ {
		preSum += nums[i]
		if c, ok := preSumFreq[preSum-k]; ok {
			result += c
		}
		preSumFreq[preSum] += 1
	}

	return result
}
func main()  {
	data := []int{1,1,1}

	fmt.Println(subarraySum2(data, 2))
}
