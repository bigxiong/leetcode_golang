package main

func numIslands(grid [][]byte) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		grid[i][j] = '0'
		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(grid, i-1, j)
		}
		if i+1 < rowLen && grid[i+1][j] == '1' {
			dfs(grid, i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(grid, i, j-1)
		}
		if j+1 < colLen && grid[i][j+1] == '1' {
			dfs(grid, i, j+1)
		}
	}

	var numberOfIslands int
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == '1' {
				numberOfIslands++
				dfs(grid, i, j)
			}
		}
	}
	return numberOfIslands
}

// BFS
func numIslands2(grid [][]byte) int {
	if grid == nil {
		return 0
	}

	rowLen := len(grid)
	colLen := len(grid[0])
	type pair struct {
		x, y int
	}
	var numberOfIslands int = 0

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if grid[i][j] == '1' {
				numberOfIslands++
				// BFS
				grid[i][j] = '0'
				var queue []pair
				queue = append(queue, pair{
					x: i,
					y: j,
				})
				for len(queue) > 0 {
					// pop from queue
					front := queue[0]
					queue = queue[1:]

					if front.x-1 >= 0 && grid[front.x-1][front.y] == '1' {
						queue = append(queue, pair{front.x - 1, front.y})
						// 标记为0， 避免重复访问
						grid[front.x-1][front.y] = '0'
					}
					if front.x+1 < rowLen && grid[front.x+1][front.y] == '1' {
						queue = append(queue, pair{front.x + 1, front.y})
						// 标记为0， 避免重复访问
						grid[front.x+1][front.y] = '0'
					}
					if front.y-1 >= 0 && grid[front.x][front.y-1] == '1' {
						queue = append(queue, pair{front.x, front.y - 1})
						// 标记为0， 避免重复访问
						grid[front.x][front.y-1] = '0'
					}
					if front.y+1 < colLen && grid[front.x][front.y+1] == '1' {
						queue = append(queue, pair{front.x, front.y + 1})
						// 标记为0， 避免重复访问
						grid[front.x][front.y+1] = '0'
					}
				}
			}
		}
	}
	return numberOfIslands
}

func main() {
	numIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	})
}
