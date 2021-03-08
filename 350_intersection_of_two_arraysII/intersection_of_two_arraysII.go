package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	helperMap := make(map[int]int)
	for _, num := range nums1 {
		helperMap[num] += 1
	}
	var k int
	for _, num := range nums2 {
		if  helperMap[num] > 0 {
			helperMap[num] -= 1
			nums2[k] = num
			k++
		}
	}
	return nums2[:k]
}

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))

	var ans []int
	var i,j int
	for i < len(nums1) && j < len(nums2){
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ans = append(ans, nums1[i])
			i++
			j++
		}
	}
	return ans
}


func main() {
	
}
