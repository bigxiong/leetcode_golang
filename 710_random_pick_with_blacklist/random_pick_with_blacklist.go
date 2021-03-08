package main

type Solution struct {
	data []int

}


func Constructor(N int, blacklist []int) Solution {
	w := make(map[int]bool)
	for i := 0 ;i < N; i++{
		w[i] = true
	}
	for _, b := range blacklist {
		delete(w, b)
	}
	for k, v := range w {

	}


}


func (this *Solution) Pick() int {

}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(N, blacklist);
 * param_1 := obj.Pick();
 */

