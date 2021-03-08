package main

type MyQueue struct {
	pushStack []int
	popStack  []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.pushStack = append(this.pushStack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			// pop
			top := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]

			this.popStack = append(this.popStack, top)
		}
	}
	if len(this.popStack) == 0 {
		return -1
	}
	top := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return top
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) > 0 {
			// pop
			top := this.pushStack[len(this.pushStack)-1]
			this.pushStack = this.pushStack[:len(this.pushStack)-1]

			this.popStack = append(this.popStack, top)
		}
	}
	if len(this.popStack) == 0 {
		return -1
	}
	top := this.popStack[len(this.popStack)-1]
	return top
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.popStack) == 0  && len(this.pushStack) == 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

