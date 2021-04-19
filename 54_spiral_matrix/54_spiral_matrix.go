package main

import "fmt"

// 顺时针打印矩阵
func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	m := len(matrix)
	n := len(matrix[0])
	var res []int
	visisted := make(map[string]bool)
	var dfs func(matrix [][]int, i, j, step, xStart, xEnd, yStart, yEnd int)
	dfs = func(matrix [][]int, i, j, step, xStart, xEnd, yStart, yEnd int) {
		if i < xStart || i > xEnd {
			return
		}
		if j < yStart || j > yEnd {
			return
		}
		key := fmt.Sprintf("%d_%d", i, j)
		if _, ok := visisted[key]; ok {
			return
		}
		res = append(res, matrix[i][j])
		visisted[key] = true

		if i == xStart && j <= yEnd {
			dfs(matrix, i, j+1, step, xStart, xEnd, yStart, yEnd)
		}

		if j == yEnd && i <= xEnd {
			dfs(matrix, i+1, j, step, xStart, xEnd, yStart, yEnd)
		}

		if i == xEnd && j > yStart {
			dfs(matrix, i, j-1, step, xStart, xEnd, yStart, yEnd)
		}

		if j == yStart && i >= xStart {
			dfs(matrix, i-1, j, step, xStart, xEnd, yStart, yEnd)
		}
	}

	x1, x2, y1, y2 := 0, m-1, 0, n-1
	step := 0
	for x1 <= x2 || y1 <= y2 {
		dfs(matrix, x1, y1, step, x1, x2, y1, y2)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(matrix[i][j], " ")
				if j == n-1 {
					fmt.Println()
				}
			}
		}
		fmt.Println()
		x1++
		x2--
		if x1 > x2 {
			break
		}
		y1++
		y2--
		if y1 > y2 {
			break
		}
	}
	return res
}

// 打印蛇形矩阵(逆时针打印)
func generateSnakeMatrix(m, n int) {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	var dfs func(matrix [][]int, i, j, step, xStart, xEnd, yStart, yEnd int) int
	dfs = func(matrix [][]int, i, j, step, xStart, xEnd, yStart, yEnd int) int {
		if i < xStart || i > xEnd {
			return step
		}
		if j < yStart || j > yEnd {
			return step
		}

		if matrix[i][j] != 0 {
			return step
		}

		step++
		if matrix[i][j] == 0 {
			matrix[i][j] = step
		}

		if i <= xEnd && j == yStart {
			step = dfs(matrix, i+1, j, step, xStart, xEnd, yStart, yEnd)
		}

		if i == xEnd && j <= yEnd {
			step = dfs(matrix, i, j+1, step, xStart, xEnd, yStart, yEnd)
		}

		if i >= xStart && j == yEnd {
			step = dfs(matrix, i-1, j, step, xStart, xEnd, yStart, yEnd)
		}

		if i == xStart && j >= yStart {
			step = dfs(matrix, i, j-1, step, xStart, xEnd, yStart, yEnd)
		}
		return step
	}

	x1, x2, y1, y2 := 0, m-1, 0, n-1
	step := 0
	for x1 <= x2 || y1 <= y2 {
		step = dfs(matrix, x1, y1, step, x1, x2, y1, y2)
		printMatrix(matrix)
		x1++
		x2--
		y1++
		y2--
	}
}

func printMatrix(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
			if j == n-1 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
}

func main() {

}
