package main

import "fmt"

/*

Given two strings s and t which consist of only lowercase letters.

String t is generated by random shuffling string s and then add one more letter at a random position.

Find the letter that was added in t.

*/

func findTheDifference(s string, t string) byte {
	helper := make(map[byte]int)

	for i := 0; i < len(t); i++{
		helper[t[i]]++
	}
	for i := 0; i < len(s); i++ {
		helper[s[i]]--
	}

	var diff byte
	for k, v := range helper {
		if v == 1 {
			diff = k
		}
	}
	return diff
}

func main() {
	fmt.Println(findTheDifference("a", "aa"))
	fmt.Println(110^110^987)
}