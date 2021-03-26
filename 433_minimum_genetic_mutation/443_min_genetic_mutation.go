package main

import "fmt"

func minMutation(start string, end string, bank []string) int {
	bankHashMap := make(map[string]bool)
	for _, b := range bank {
		bankHashMap[b] = true
	}
	if start == end && bankHashMap[start] {
		return 0
	}

	geneChoices := []byte{'A', 'G', 'C', 'T'}
	var queue []string
	visited := make(map[string]bool)
	queue = append(queue, start)

	var step int = 0
	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			// pop from queue head
			item := queue[0]
			queue = queue[1:]
			if bankHashMap[item] && item == end {
				return step
			}
			// 依次变换8个字母
			for i := 0; i < 8; i++ {
				// 每个位置可以变换出3个来
				for _, geneSelected := range geneChoices {
					if geneSelected != item[i] {
						thisGene := []byte(item)
						thisGene[i] = geneSelected
						fmt.Println(item, " ", i, " ", item[i], "->", geneSelected, string(thisGene))
						if !visited[string(thisGene)] && bankHashMap[string(thisGene)] {
							queue = append(queue, string(thisGene))
							visited[string(thisGene)] = true
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
	res := minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"})
	res1 := minMutation("AACCGGTT", "AACCGGTA", nil)
	fmt.Println(res, res1)
}
