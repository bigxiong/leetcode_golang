package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	var spaceCount int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}

	strArray := make([]byte, len(s)+2*spaceCount)
	var k int = len(strArray) - 1
	for i := len(s) - 1; i >= 0; i-- {
		// %20
		if s[i] == ' ' {
			strArray[k] = '0'
			strArray[k-1] = '2'
			strArray[k-2] = '%'
			k = k - 2
		} else {
			strArray[k] = s[i]
			k--
		}
	}
	return string(strArray)
}

func replaceSpace2(s string) string {
	var strArray []rune
	for _, c := range s {
		if c == ' ' {
			strArray = append(strArray, '%', '2', '0')
		} else {
			strArray = append(strArray, c)
		}
	}
	return string(strArray)
}

func replaceSpace3(s string) string {
	var res strings.Builder
	for _, c := range s {
		if c == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteRune(c)
		}
	}
	return res.String()
}

func main() {
	str := replaceSpace3("We are happy.")
	fmt.Println(str)
}
