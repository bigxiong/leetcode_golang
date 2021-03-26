package main

import (
	"fmt"
	"sync"
)

func main() {
	//var wg sync.WaitGroup
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//go func(){
	//	for  i := 1; i <= 26; i++ {
	//		<- ch1
	//		fmt.Printf("%d",i)
	//		ch2 <- i
	//	}
	//	wg.Done()
	//	ch2 <- 26
	//}()
	//
	//go func(){
	//	for  i := 'A'; i <= 'Z'; i++ {
	//		<- ch2
	//		fmt.Printf("%c",i)
	//		if i != 'Z' {
	//			ch1 <- 1
	//		}
	//	}
	//	wg.Done()
	//}()
	//ch1 <- 1
	//wg.Add(2)
	//wg.Wait()
	wg := &sync.WaitGroup{}
	wg.Add(2)

	limit := 26

	numChan := make(chan int, 1)
	charChan := make(chan int, 1)
	charChan <- 1

	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			<-charChan
			fmt.Printf("%c\n", 'a'+i)
			numChan <- 1

		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			<-numChan
			fmt.Println(i)
			charChan <- 1

		}
	}()

	wg.Wait()

	close(charChan)
	close(numChan)
}
