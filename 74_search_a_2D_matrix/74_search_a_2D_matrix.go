package main

// 注意此题与剑指offer04题目的比较
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	// 二分查找
	left := 0
	right := m*n - 1
	var pivotIdx, pivotElement int
	for left <= right {
		pivotIdx = (left + right) / 2
		pivotElement = matrix[pivotIdx/n][pivotIdx%n]
		if target == pivotElement {
			return true
		} else if target < pivotElement {
			right = pivotIdx - 1
		} else {
			left = pivotIdx + 1
		}
	}
	return false
}

func main() {

}
