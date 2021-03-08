package main

import "fmt"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func removeDuplicateLetters(s string) string {
	var inStack [256]int
	var stack []byte
	var letterFreq [256]int
	for i := 0; i < len(s); i++ {
		letterFreq[s[i]]++
	}
	for i := 0 ; i < len(s); i++ {
		c := s[i]
		// 遍历过该字符，就减1
		letterFreq[s[i]]--

		// 当前字符不在栈中
		if inStack[c] == 0 {
			for len(stack) > 0 && stack[len(stack)-1] > c {
				if letterFreq[stack[len(stack)-1]] == 0 {
					break
				}
				inStack[stack[len(stack)-1]] = 0
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, c)
			inStack[c] = 1
			// 放在这个地方错误的
			// letterFreq[s[i]]--
		}
	}

	var sArray []byte
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sArray = append(sArray, top)
	}


	return Reverse(string(sArray))
}

func main()  {
	//fmt.Println(removeDuplicateLetters("bcabc"))

	fmt.Println(removeDuplicateLetters("bbcaac"))
}