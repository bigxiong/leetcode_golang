package main

import (
	"strings"
)

func reverseWords(s string) string {
	if s == "" {
		return s
	}
	runeStr := []rune(s)
	var stringBuilder []string
	var k int
	for k < len(runeStr) && runeStr[k] == 32 {
		k++
	}
	for i := k; i < len(runeStr); i++ {
		if runeStr[i] == 32 || i == len(runeStr)-1{
			if i == len(runeStr)-1 && runeStr[i] != 32{
				stringBuilder = append(stringBuilder, string(runeStr[k:]))
			} else {
				stringBuilder = append(stringBuilder, string(runeStr[k:i]))
			}
			for i < len(runeStr)-1 && runeStr[i] == 32{
				i++
			}
			k = i
		}
	}
	for i,j := 0, len(stringBuilder)-1; i < j; i,j = i+1, j-1 {
		stringBuilder[i], stringBuilder[j] = stringBuilder[j], stringBuilder[i]
	}
	return strings.Join(stringBuilder," ")
}

func main()  {
	//input1 := "the sky is blue"
	////          "blue is sky the"
	//input2 := "  hello world!  "
	//input3 := "a good   example"

	//fmt.Println(reverseWords(input1))
	//fmt.Println(reverseWords(input2))
	//fmt.Println(reverseWords(input3))
	reverseWords("1 ")
	//fmt.Printf("%q\n", input1)
	//fmt.Printf("%+q\n", input1)
	//rune1 := []rune(input1)


	//
	//fmt.Printf("before: %v, after: %v \n", input1, reverseWords(input1))
	//fmt.Printf("before: %v, after: %v \n", input2, reverseWords(input2))
	//fmt.Printf("before: %v, after: %v \n", input3, reverseWords(input3))

	//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	//fmt.Println(sample)
	//for i := 0; i < len(sample); i++ {
	//	fmt.Printf("%x ", sample[i])
	//}
	//
	//fmt.Printf("\n% x\n", sample)
	//fmt.Printf("%q\n", sample)
	//fmt.Printf("%+q\n", sample)
	//
	//
	//const placeOfInterest = `⌘`
	//
	//fmt.Printf("plain string: ")
	//fmt.Printf("%s", placeOfInterest)
	//fmt.Printf("\n")
	//
	//fmt.Printf("quoted string: ")
	//fmt.Printf("%+q", placeOfInterest)
	//fmt.Printf("\n")
	//
	//fmt.Printf("hex bytes: ")
	//for i := 0; i < len(placeOfInterest); i++ {
	//	fmt.Printf("%x ", placeOfInterest[i])
	//}
	//fmt.Printf("\n")
}
