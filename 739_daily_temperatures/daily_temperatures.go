package main

import "fmt"

func dailyTemperatures(T []int) []int {
	var res = make([]int, len(T))
	type entry struct {
		index int
		val   int
	}

	var stack []entry
	for i := len(T) - 1 ; i >= 0 ; i-- {
		for len(stack) > 0 && T[i] > stack[len(stack)-1].val {
			// pop
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			res[i] = top.index - i
		} else {
			res[i] = 0
		}
		stack = append(stack, entry{
			index: i,
			val:   T[i],
		})
	}
	return res
}

func main()  {
	d := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(d))
}