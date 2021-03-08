package main

import (
	"container/list"
	"fmt"
)

type LFUNode2 struct {
	key int
	val int
	freq int
}

type LFUCache2 struct {
	cache                   map[int]*list.Element
	freq                    map[int]*list.List
	capacity, size, minFreq int
}

func NewLFUCache2(capacity int) LFUCache2 {
	return LFUCache2 {
		cache:    make(map[int]*list.Element),
		freq:     make(map[int]*list.List),
		capacity: capacity,
	}
}

func (this *LFUCache2) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.IncFreq(node)
		return node.Value.(LFUNode2).val
	}
	return -1
}

func (this *LFUCache2) Put(key, value int) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.cache[key]; ok {
		node.Value = LFUNode2{
			key:  key,
			val:  value,
			freq: node.Value.(LFUNode2).freq,
		}
		this.IncFreq(node)
	} else {
		if len(this.cache) >= this.capacity {
			removeE := this.freq[this.minFreq].Back()
			// 淘汰一个
			node := this.freq[this.minFreq].Remove(removeE)
			delete(this.cache, node.(LFUNode2).key)
		}

		x := LFUNode2{key: key, val: value, freq: 1}
		if this.freq[x.freq] == nil {
			this.freq[1] = list.New()
		}
		xElement := this.freq[1].PushFront(x)
		this.cache[key] = xElement
		this.minFreq = 1
	}
}

// 元素被访问后，更新它的频次
func (this *LFUCache2) IncFreq(element *list.Element){
	node := element.Value.(LFUNode2)
	_freq := node.freq
	// 1. 从该freq 列表中移除该节点
	this.freq[_freq].Remove(element)

	// _freq对应的链表为minFreq, 且只有一个节点
	if this.minFreq == _freq && this.freq[_freq].Len() == 0 {
		// 最小freq 为_freq_1
		this.minFreq = _freq+1
		// 从freq map中移除
		delete(this.freq, _freq)
	}

	node.freq++
	if _, ok := this.freq[node.freq]; !ok {
		this.freq[node.freq] = list.New()
	}
	// 活跃的元素在list 的Front
	this.freq[node.freq].PushFront(element)
}

func main() {
	obj := NewLFUCache2(2)
	obj.Put(3,1) // (3,1,1,1)
	obj.Put(2,1) // (3,1,1,1), (2,1,1,2)
	obj.Put(2,2) // (3,1,1,1), (2,2,2,3)
	obj.Put(4,4) // (4,4,1,4), (2,1,2,3)
	fmt.Println(obj.Get(2))

	obj = NewLFUCache2(0)
	obj.Put(0,0) // (3,1,1,1)
	obj.Get(0)
}
