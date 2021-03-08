package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	data []int
	hashTable map[int]int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		data:      nil,
		hashTable: make(map[int]int),
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashTable[val]; ok {
		return false
	}
	this.data = append(this.data, val)
	this.hashTable[val] = len(this.data)-1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hashTable[val]; !ok {
		return false
	}
	index := this.hashTable[val]
	// delete val->index mapping
	delete(this.hashTable, val)
	tail := len(this.data)-1
    if index == tail {
		// remove tail
		this.data = this.data[:tail]
		return true
	}

	// swap index <-> tail
	this.data[index] = this.data[tail]
	// update index
	this.hashTable[this.data[index]] = index
	// remove tail
	this.data = this.data[:tail]
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.data))
	return this.data[random]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main()  {
	/*
	["RandomizedSet","remove","remove","insert","getRandom","remove","insert"]
	[[],[0],[0],[0],[],[0],[0]]
	*/

	obj := Constructor()
	param_1 := obj.Remove(0)
	param_2 := obj.Remove(0)
	param_3 := obj.Insert(0)
	param_4 := obj.GetRandom()
	param_5 := obj.Remove(0)
	param_6 := obj.Insert(0)

	fmt.Println(param_1)
	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)

}