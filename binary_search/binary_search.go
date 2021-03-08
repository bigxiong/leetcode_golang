package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}


func BinarySearch(a []int, v int) int {
	left := 0
	right := len(a)-1
	for left <= right {
		mid := left + (right - left)/2
		if v == a[mid] {
			return mid
		} else if v < a[mid] {
			right = mid - 1
		} else if a[mid] < v {
			left = mid + 1
		}
	}
	return -1
}

func BinarySearchFirst(a []int, v int) int {
	left := 0
	right := len(a)-1
	for left <= right {
		mid := left + (right - left)/2
		if v < a[mid] {
			right = mid - 1
		} else if a[mid] < v {
			left = mid + 1
		} else {
			if mid == 0 || a[mid-1] != v {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func BinarySearchLast(a []int, v int) int {
	left := 0
	right := len(a)-1
	for left <= right {
		mid := left + (right - left)/2
		if v < a[mid] {
			right = mid - 1
		} else if a[mid] < v {
			left = mid + 1
		} else {
			if mid == len(a)-1 || a[mid+1] != v {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func BinarySearchGT(a []int, v int) int {
	left := 0
	right := len(a)-1
	for left <= right {
		mid := left + (right - left)/2
		if v < a[mid] {
			right = mid - 1
		} else if a[mid] < v {
			left = mid + 1
		} else {
			if mid != len(a)-1 && a[mid+1] > v {
				return mid+1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
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
	mySqrt(5)
	a := []int{1,2,3,3,3,3,4,5,6}
	fmt.Println(BinarySearchLT(a, 3))
	fmt.Println(BinarySearchGT(a, 3))
}
