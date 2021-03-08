package main

import "fmt"

func fibonacci(i int) int {
	if i <= 1 {
		return i
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func FibonacciLoop(n int) int {
	if n <= 1 {
		return n
	}
	f1 := 0
	f2 := 1
	f3 := f1 + f2
	for i := 2; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}
func main()  {
	for i := 0; i < 20; i++ {
		//fmt.Println(fibonacci(i))
		fmt.Println(FibonacciLoop(i))
	}
}