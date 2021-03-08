package main

import (
	"fmt"
	"sort"
)

/*

Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.

Please note that the string does not contain any non-printable characters.

Example:

Input: "Hello, my name is John"
Output: 5

*/

func countSegments(s string) int {
	runeStr := []rune(s)
	var segmentsCount int
	var i int
	// Skip spaces in the front of the input.
	for i < len(runeStr) && runeStr[i] == 32 {
		i++
	}
	//fieldStart := i
	for i < len(runeStr) {
		for i < len(runeStr)  && runeStr[i] != 32 {
			i++
		}
		segmentsCount++
		for i < len(runeStr) && runeStr[i] == 32 {
			i++
		}
		//fieldStart = i
	}

	return segmentsCount

}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
	fmt.Println(countSegments("                "))
	//countSegments("Hello, my name is John")


}
