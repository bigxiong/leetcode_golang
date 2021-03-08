package main

import (
	"fmt"
)

// 1. Quick Sort
func quickSort(data []int, left int, right int)  {
	if left >= right {
		return
	}
	p := partition2(data, left, right)
	quickSort(data, left, p-1)
	quickSort(data, p+1, right)
}

func partition1(data []int, left, right int) int {
	pivot := data[left]
	i := left
	j := right
	// i's job : index <= i 的元素都<= pivot
	for ; i < j; {
		// 找到右侧第一个大于pivot的元素
		for i < j && data[j] > pivot {
			j--
		}
		data[i] = data[j]
		for i < j && data[i] <= pivot {
			i++
		}
		data[j] = data[i]
	}
	data[i] = pivot
	return i
}

func partition2(data []int, left, right int) int {
	pivot := data[left]
	//pivotIdx := left
	i := left
	j := right
	for ; i < j; {
		for i < j && data[j] <= pivot {
			j--
		}
		data[i] = data[j]
		for i < j && data[i] >= pivot {
			i++
		}
		//data[i],data[j] = data[j],data[i]
		data[j] = data[i]
	}
	//data[pivotIdx], data[i] = data[i], data[pivotIdx]
	data[i] = pivot
	return i
}

// 2. Merge Sort
func mergeSort(data []int, left int, right int)  {
	if left >= right {
		return
	}
	mid := left + (right - left) / 2
	mergeSort(data, left, mid)
	mergeSort(data, mid + 1, right)

	merge(data, left, mid, right)

}

func merge(data []int, left, mid, right int)  {
	tmp := make([]int, right-left+1)

	i := left
	j := mid+1
	k := 0
	for ;i <= mid && j <= right; k++ {
		if data[i] <= data[j] {
			tmp[k] = data[i]
			i++
		} else {
			tmp[k] = data[j]
			j++
		}
	}

	for i <= mid {
		tmp[k] = data[i]
		k++
		i++
	}
	for j <= right {
		tmp[k] = data[j]
		k++
		j++
	}
	copy(data[left:right+1], tmp)

}

// 3. Kth Largest Element in an Array
/*
Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
*/
func findKthLargest(nums []int, k int) int {

	left := 0
	right := len(nums) - 1
	for {
		p := partition2(nums, left, right)
		if p+1 == k {
			return nums[p]
		} else if p+1 < k {
			left = p+1
		} else {
			right = p-1
		}
	}
}

// 4. Insertion sort
func insertionSort(nums []int)  {
	for i := 1; i < len(nums); i++{
		tmp := nums[i]
		j := i-1
		for ; tmp < nums[j] && j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmp
	}
}

// 5. counting sort 计数排序
func countingSort(nums []int)  {
	if len(nums) <= 0 {
		return
	}
	// 建立一个计数数组
	var min = nums[0]
	var max = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	countArrayLen := max-min+1
	countArray := make([]int, countArrayLen)
	// 统计数组
	for i := 0; i < len(nums); i++ {
		countArray[nums[i]-min]++
	}
	// 扫描输出
	var index int
	for i := 0; i < len(countArray); i++{
		num := i
		count := countArray[i]
		for count > 0 {
			nums[index] = num+min
			index++
			count--
		}
	}
}


func countingSort2(nums []int) []int {
	if len(nums) <= 0 {
		return nil
	}
	// 建立一个计数数组
	var min = nums[0]
	var max = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	countArrayLen := max-min+1
	countArray := make([]int, countArrayLen)
	// 统计数组
	for i := 0; i < len(nums); i++ {
		countArray[nums[i]-min]++
	}

	//累加数组
	var sum int
	for i, c := range countArray {
		sum += c
		countArray[i] = sum
	}


	// 扫描输出
	sortedArray := make([]int, len(nums))
	for j := len(nums)-1; j >= 0 ; j--{
		num := nums[j]
		countIndex := nums[j]-min
		index := countArray[countIndex]
		sortedArray[index-1] = num
		countArray[countIndex]--
	}
	return sortedArray
}


// 6. bucket sort
// [0, 99]
func bucketSort(nums []int)  {
	// 10个桶, 每个桶是一个链
	var min = nums[0]
	var max = nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	d := max - min
	bucketNum := len(nums)

	bucket := make([][]int, bucketNum)

	for i := 0; i < len(nums); i++{
		num := nums[i]
		bucketIdx := (num-min)*(bucketNum-1) / d
		bucketE := bucket[bucketIdx]
		bucketE = append(bucketE, num)
		// insertion sort. stable
		for j := len(bucketE)-1; j > 0 ; j-- {
			if bucketE[j-1] > nums[i] {
				bucketE[j] = bucketE[j-1]
			} else {
				bucketE[j] = nums[i]
				break
			}
		}

		bucket[bucketIdx] = bucketE
	}
	var index int
	for i := 0; i < len(bucket); i++ {
		b := bucket[i]
		for _, e := range b {
			nums[index] = e
			index++
		}

	}

}

func bubbleSort(nums []int)  {
	var flag bool
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++{
			if nums[j] > nums[j+1] {
				nums[j],nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}


}

func main() {
	//data := []int{3,4,8,2,5,6,1,9,7}
	////mergeSort(data, 0, len(data)-1)
	////fmt.Println(data)
	//fmt.Println(data)
	//quickSort(data, 0, len(data)-1)
	//fmt.Println(data)

	data := []int{12,13,24,12,25,16,17,19,1}

	quickSort(data, 0, len(data)-1)
	fmt.Println(data)


	kth := findKthLargest([]int{3,2,3,1,2,4,5,6}, 3)

	fmt.Println(kth)
}

