package main

import "fmt"

//func exist(board [][]byte, word string) bool {
//	row, col := len(board), len(board[0])
//	var visited = make([][]bool, len(board))
//	for i := 0; i < len(board); i++ {
//		visited[i] = make([]bool, len(board[0]))
//	}
//	var backtrack func(board [][]byte, x, y int, word string, selected []byte) bool
//
//	backtrack = func(board [][]byte, x, y int, word string, selected []byte) bool{
//		if word == string(selected) {
//			return true
//		}
//
//		if x - 1 >= 0 && !visited[x-1][y] {
//			visited[x-1][y] = true
//			selected = append(selected, board[x-1][y])
//			isOk := backtrack(board, x-1, y, word, selected)
//			selected = selected[:len(selected)-1]
//			visited[x-1][y] = false
//			if isOk {
//				return true
//			}
//		}
//
//		if x + 1 < len(board) && !visited[x+1][y]{
//			visited[x+1][y] = true
//			selected = append(selected, board[x+1][y])
//			isOk := backtrack(board, x+1, y, word, selected)
//			selected = selected[:len(selected)-1]
//			visited[x+1][y] = false
//			if isOk {
//				return true
//			}
//		}
//
//		if y - 1 >= 0 && !visited[x][y-1] {
//			visited[x][y-1] = true
//			selected = append(selected, board[x][y-1])
//			isOk := backtrack(board, x, y-1, word, selected)
//			selected = selected[:len(selected)-1]
//			visited[x][y-1] = false
//			if isOk {
//				return true
//			}
//		}
//
//		if y + 1 < len(board[0]) && !visited[x][y+1]{
//			visited[x][y+1] = true
//			selected = append(selected, board[x][y+1])
//			isOk := backtrack(board, x, y+1, word, selected)
//			selected = selected[:len(selected)-1]
//			visited[x][y+1] = false
//			if isOk {
//				return true
//			}
//		}
//
//
//		return false
//	}
//
//
//	for i := 0; i < row; i++ {
//		for j := 0; j < col; j++ {
//			var selected []byte
//			selected = append(selected, board[i][j])
//			visited[i][j] = true
//			isOk :=  backtrack(board, i, j, word, selected)
//			if isOk {
//				return true
//			}
//			visited[i][j] = false
//		}
//	}
//
//	return false
//}

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var backtrack func(board [][]byte, x, y int, word string) bool

	backtrack = func(board [][]byte, x, y int, word string) bool {
		// 1. 匹配完成
		if len(word) == 0 {
			return true
		}
		// 2.
		if board[x][y] != word[0] {
			return false
		}
		// 3.
		if len(word) == 1 {
			return true
		}

		visited[x][y] = true
		// 上
		if x-1 >= 0 && !visited[x-1][y] {
			isOk := backtrack(board, x-1, y, word[1:])
			if isOk {
				return true
			}
		}

		// 下
		if x+1 < len(board) && !visited[x+1][y] {
			isOk := backtrack(board, x+1, y, word[1:])
			if isOk {
				return true
			}
		}

		if y-1 >= 0 && !visited[x][y-1] {
			isOk := backtrack(board, x, y-1, word[1:])
			if isOk {
				return true
			}
		}

		if y+1 < len(board[0]) && !visited[x][y+1] {
			isOk := backtrack(board, x, y+1, word[1:])
			if isOk {
				return true
			}
		}

		visited[x][y] = false
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				isOk := backtrack(board, i, j, word)
				if isOk {
					return true
				}
			}
		}
	}

	return false
}

func exist2(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	var backtrack func(board [][]byte, x, y int, word string) bool
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	backtrack = func(board [][]byte, x, y int, word string) bool {
		// 1. 匹配完成
		if len(word) == 0 {
			return true
		}
		// 2.
		if board[x][y] != word[0] {
			return false
		}
		// 3.
		if len(word) == 1 {
			return true
		}

		visited[x][y] = true
		// 上
		//if x - 1 >= 0 && !visited[x-1][y] {
		//	isOk := backtrack(board, x-1, y, word[1:])
		//	if isOk {
		//		return true
		//	}
		//}
		//
		//// 下
		//if x + 1 < len(board) && !visited[x+1][y]{
		//	isOk := backtrack(board, x+1, y, word[1:])
		//	if isOk {
		//		return true
		//	}
		//}
		//
		//if y - 1 >= 0 && !visited[x][y-1] {
		//	isOk := backtrack(board, x, y-1, word[1:])
		//	if isOk {
		//		return true
		//	}
		//}
		//
		//if y + 1 < len(board[0]) && !visited[x][y+1]{
		//	isOk := backtrack(board, x, y+1, word[1:])
		//	if isOk {
		//		return true
		//	}
		//}
		for _, direction := range directions {
			x := x + direction[0]
			y := y + direction[1]
			if 0 <= x && x < row && 0 <= y && y < col && !visited[x][y] {
				isOk := backtrack(board, x, y, word[1:])
				if isOk {
					return true
				}
			}
		}

		visited[x][y] = false
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] {
				isOk := backtrack(board, i, j, word)
				if isOk {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	//board := [][]byte{
	//	{'A','B','C','E'},
	//	{'S','F','C','S'},
	//	{'A','D','E','E'},
	//}
	board := [][]byte{
		{'A', 'A'},
	}

	b := exist(board, "AA")
	fmt.Println(b)
}
