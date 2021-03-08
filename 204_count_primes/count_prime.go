package main

import "fmt"

func countPrimes(n int) int {
	isPrim := make([]bool, n)
	for i := 0 ; i < len(isPrim); i++{
		isPrim[i] = true
	}
	for i := 2; i*i < n; i++{
		if isPrim[i] {
			for j := 2*i; j < n;j += i {
				isPrim[j] = false
			}
		}
	}

	var count int
	for i := 0; i < len(isPrim); i++{
		if isPrim[i] {
			count++
		}
	}
	return count
}


// A concurrent prime sieve

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <- ch
		fmt.Println(prime)
		out := make(chan int)
		go Filter(ch, out, prime)
		ch = out
	}
}