package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if target < matrix[i][j] {
			j--
		} else {
			i++
		}
	}
	return false
}

func main() {

}
