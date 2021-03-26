package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	numberOfCookies := len(s)
	numberOfChild := len(g)
	var ans int
	for i, j := 0, 0; i < numberOfChild && j < numberOfCookies; i++ {
		// cookie 不能满足该小孩
		for j < numberOfCookies && g[i] > s[j] {
			j++
		}
		if j < numberOfCookies {
			ans++
			j++
		}
	}
	return ans
}

func main() {

}
