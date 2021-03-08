package main

import (
	"fmt"
	"sort"
)

/**
two sum:
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
。
*/
// method 2: Sort + two pointer
// O(nlogn)
func twoSum2(nums []int, target int) []int {
	sort.Ints(nums)
	for i, j := 0, len(nums)-1 ; i < j; {
		if target == nums[i] + nums[j] {
			return []int{i,j}
		} else if target < nums[i] + nums[j] {
			j--
		} else {
			i++
		}
	}
	return nil
}

// method 2: hashmap
// O(n)
func twoSum1(nums []int, target int) []int {
	hashTable := make(map[int]int)
	for i, x := range nums {
		if j, ok := hashTable[target-x]; ok {
			return []int{j, i}
		}
		hashTable[x] = i
	}
	return nil
}

// method 3: two sum 去掉重复部分
func twoSum3(nums []int, start int, target int) [][]int  {
	var result [][]int
	sort.Ints(nums)
	left := start
	right := len(nums)-1
	for left < right {
		if  nums[left] + nums[right] < target {
			left++
			for left < right && nums[left-1] == nums[left] {
				left++
			}
		} else if target < nums[left] + nums[right] {
			right--
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else {
			result = append(result, []int{nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left-1] == nums[left] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		x1 := nums[i]
		tuples := twoSum3(nums, i+1, 0-x1)
		for _, tuple := range tuples{
			result = append(result, []int{x1, tuple[0], tuple[1]})
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums); i++{
		// 跳过重复的元素 使用if ... continue
		// 不要使用for ... i++ ,
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		x1 := nums[i]
		// two sum问题
		left, right := i+1, len(nums) - 1
		for left < right {
			x2 := nums[left]
			x3 := nums[right]
			if x2 + x3 + x1 == 0 {
				result = append(result, []int{x1, x2, x3})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if x2 + x3 + x1  < 0 {
				left++
			} else{
				right--
			}
		}
	}
	return result
}


func fourSum(nums []int, target int) [][]int {
	var resultSet map[[4]int]bool
	resultSet = make(map[[4]int]bool)
	sort.Ints(nums)
	hashTable := make(map[int][][]int)
	for i := 0; i < len(nums)-3; i++{
		for j := i+1; j < len(nums)-2; j++{
			currSum := nums[i] + nums[j]
			var curPairs [][]int
			curPairs = hashTable[currSum]
			curPairs = append(curPairs, []int{i, j})
			hashTable[currSum] = curPairs
		}
	}
	for i := 1; i < len(nums)-1; i++{
		for j := i+1; j < len(nums); j++{
			currSum := nums[i] + nums[j]
			if prevPairs, ok := hashTable[target-currSum]; ok {
				for _, pair := range prevPairs {
					if pair[1] < i {
						r := [4]int{nums[pair[0]], nums[pair[1]], nums[i], nums[j]}
						fmt.Println(r)
						resultSet[r] = true
					}
				}

			}
		}
	}
	var result [][]int
	for k,_ := range resultSet {
		fmt.Println("result set: ", k)
		var r []int
		r = append(r, k[0], k[1] , k[2], k[3])
		result = append(result, r)
	}
	fmt.Println("result: ", result)
	return  result
}

func main() {
	fourSum([]int{1,0,-1,0,-2,2}, 0)
	threeSum1([]int{1,-1,-1,0})
}
