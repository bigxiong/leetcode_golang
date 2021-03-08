package main

import (
	"container/heap"
	"fmt"
	"time"
)

type LFUNode struct {
	index int
	key  int
	val  int
	freq int
	tick int64
}

type PriorityQueue []*LFUNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 1. freq频率低，优先级小
	if pq[i].freq < pq[j].freq {
		return true
	// 2. freq频率一样，tick越小(最近访问时间最久远), 优先级越小
	} else if pq[i].freq == pq[j].freq {
		return pq[i].tick < pq[j].tick
	} else {
		return false
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*LFUNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *LFUNode) {
	item.freq++
	item.tick = time.Now().UnixNano()
	heap.Fix(pq, item.index)
}

type LFUCache struct {
	capacity  int
	items 	  map[int]*LFUNode
	freqHeap  PriorityQueue
}


func Constructor1(capacity int) LFUCache {
	return LFUCache{
		items:    make(map[int]*LFUNode),
		capacity: capacity,
	}
}


func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	item, isExsit := this.items[key]
	if !isExsit {
		return -1
	}
	this.freqHeap.update(item)
	return item.val
}


func (this *LFUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}
	item, isExsit := this.items[key]
	if isExsit {
		item.val = value
		this.freqHeap.update(item)
	} else {
		node := LFUNode{
			key:   key,
			val:   value,
			freq:  1,
			tick:  time.Now().UnixNano(),
		}
		// 根据freq, tick 淘汰掉1个
		if len(this.items) >= this.capacity {
			item := heap.Pop(&this.freqHeap).(*LFUNode)
			delete(this.items, item.key)
		}
		this.items[key] = &node
		heap.Push(&this.freqHeap, &node)
	}
}


func main() {
	obj := Constructor1(2)
	obj.Put(3,1) // (3,1,1,1)
	obj.Put(2,1) // (3,1,1,1), (2,1,1,2)
	obj.Put(2,2) // (3,1,1,1), (2,2,2,3)
	obj.Put(4,4) // (4,4,1,4), (2,1,2,3)
	fmt.Println(obj.Get(2))

	obj = Constructor1(0)
	obj.Put(0,0) // (3,1,1,1)
	obj.Get(0)
}
