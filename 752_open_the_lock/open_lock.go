package main

import "fmt"

func openLock(deadends []string, target string) int {
	var queue []string
	queue = append(queue, "0000")
	dead := make(map[string]bool)
	for _, d := range deadends{
		dead[d] = true
	}
	visited := make(map[string]bool)// 已访问过的集合，并记录到达该状态花的步数
	visited["0000"] = true
	var step int = 0
	// BFS
	for len(queue) > 0 {
		size := len(queue)
		for i := 0 ; i < size; i++{
			p := queue[0]
			queue = queue[1:]
			if _, ok := dead[p]; ok {
				continue
			}
			fmt.Println(p)
			if p == target {
				return step
			}
			for j := 0 ; j < 4; j++{
				up := plusOne(p, j)
				if _, isExist := visited[up]; !isExist {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(p, j)
				if _, isExist := visited[down]; !isExist {
					queue = append(queue, down)
					visited[down] = true
				}
			}

		}
		step++
	}
	return -1
}

func plusOne(numbers string, index int) string {
	numbersArray := []rune(numbers)
	if numbersArray[index] == '9' {
		numbersArray[index] = '0'
	} else {
		numbersArray[index] += 1
	}
	return string(numbersArray)
}

func minusOne(numbers string, index int) string {
	numbersArray := []rune(numbers)
	if numbersArray[index] == '0' {
		numbersArray[index] = '9'
	} else {
		numbersArray[index] -= 1
	}
	return string(numbersArray)
}

func main() {
	step := openLock([]string{"0201","0101","0102","1212","2002"}, "0202")
	fmt.Println(step)
}
