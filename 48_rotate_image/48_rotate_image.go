package main

// n * n
// 旋转前后，坐标对应关系
// (i,j) -> (j, n-1-i)
/*
1 2 3     7 4 1
4 5 6 --> 8 5 2
7 8 9     9 6 3

Step 1. Transpose Matrix 延正对角线翻转 swap(matrix[i][j], matrix[j][i])
1 4 7
2 5 8
3 6 9

Step 2. Flip Horizontally 列水平交换 swap(matrix[i][j], matrix[i][N-1-j])
7 4 1
8 5 2
9 6 3


*/
func rotate(matrix [][]int) {
	if matrix == nil {
		return
	}

	N := len(matrix)
	for i := 0; i < N; i++ {
		// 此处是从j开始，对角线上部分
		for j := i; j < N; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N/2; j++ {
			matrix[i][j], matrix[i][N-1-j] = matrix[i][N-1-j], matrix[i][j]
		}
	}
}

func main() {
	rotate([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}
