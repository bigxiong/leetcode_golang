package main

func solveSudoku(board [][]byte)  {
    var backtrack func(board [][]byte, r, c int) bool
	backtrack = func(board [][]byte, row, col int) bool {
		m , n := 9, 9
		if col == n {
			// 穷举到最后一列的话就换到下一行重新开始。
			return backtrack(board, row + 1, 0)
		}
		if row == m {
			// 找到一个可行解，触发 base case
			return true
		}
		// 就是对棋盘的每个位置进行穷举
		for i := row; i < m; i++{
			for j := col; j < n; j++ {
				//// 做选择
				//backtrack(board, i, j)
				//// 撤销选择
				// 如果该位置是预设的数字，不用我们操心
				if board[i][j] != '.' {
					return backtrack(board, i, j + 1)
				}
				for ch := '1'; ch <= '9'; ch++ {
					// 如果遇到不合法的数字，就跳过
					if  !isValid(board, i, j, byte(ch)) {
						continue
					}
					// 做选择
					board[i][j] = byte(ch)
					// 继续穷举下一个
					if backtrack(board, i, j + 1) {
						return true
					}
					// 撤销选择
					board[i][j] = '.'
				}
				return false
			}
		}
		return false
	}
	backtrack(board, 0 , 0)
}

// 判断 board[i][j] 是否可以填入 n
func isValid(board [][]byte, r, c int, n byte)  bool {
	for i := 0; i < 9; i++ {
		// 判断行是否存在重复
		if board[r][i] == n {
			return false
		}
		// 判断列是否存在重复
		if board[i][c] == n {
			return false
		}
		// 判断 3 x 3 方框是否存在重复
		if board[(r/3)*3 + i/3][(c/3)*3 + i%3] == n {
			return false
		}
	}
	return true
}