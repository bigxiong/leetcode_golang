package main

import (
	"fmt"
	"strings"
)


func solveNQueens(n int) [][]string {
	var board [][]string
	var solutions [][]string
	// 1. build board
	for i := 0; i < n; i++{
		var row []string = make([]string, n)
		for j := 0 ; j < len(row); j++{
			row[j] = "."
		}
		board = append(board, row)
	}
	var f func([][]string, int)
	f = func (board [][]string, row int)  {
		// done
		if row == len(board) {

			var solution []string
			for i := 0 ; i < len(board); i++{
				solution = append(solution, strings.Join(board[i],""))
			}
			fmt.Println(solution)
			solutions = append(solutions, solution)
		}
		for col := 0 ; col < len(board); col++{
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = "Q"
			f(board, row + 1)
			board[row][col] = "."
		}
	}

	f(board, 0)
	fmt.Println(solutions)
	return solutions
}


func isValid(board [][]string, row int, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 右上是否冲突
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j++
	}
	// 左上是否冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == "Q" {
			return false
		}
		i--
		j--
	}
	return true
}


func main()  {
	solveNQueens(1)
}