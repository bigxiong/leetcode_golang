package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := make(map[string]bool)
	for _, b := range wordList {
		dict[b] = true
	}
	if beginWord == endWord && dict[beginWord] {
		return 0
	}

	var queue []string
	visited := make(map[string]bool)
	queue = append(queue, beginWord)

	var step int = 1
	for len(queue) > 0 {
		// 遍历当前层
		for i := len(queue); i > 0; i-- {
			// pop from queue head
			item := queue[0]
			queue = queue[1:]

			if dict[item] && item == endWord {
				return step
			}
			// 依次变换每个字母
			for i := 0; i < len(beginWord); i++ {
				// 每个位置可以变换 a-z
				for c := 'a'; c <= 'z'; c++ {
					if c != int32(item[i]) {
						curCharArray := []byte(item)
						curCharArray[i] = byte(c)
						fmt.Println(item, " ", i, " ", item[i], "->", c, string(curCharArray))
						// 没有被访问，且在字典中。加入到下一层中
						if !visited[string(curCharArray)] && dict[string(curCharArray)] {
							queue = append(queue, string(curCharArray))
							visited[string(curCharArray)] = true
							if string(curCharArray) == endWord {
								return step + 1
							}
						}
					}
				}
			}
		}
		step++
	}

	return -1
}

func main() {
	ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "cog"})
}
