package main

import "container/list"

type lruNode struct {
	key int
	val int
}
type LRU struct {
	capacity int
	evictedList *list.List
	items map[int]*list.Element
}

func NewLRU(capacity int) *LRU {
	return &LRU{capacity: capacity}
}

func (l *LRU) Get(key int) int  {
	item, ok := l.items[key]
	if ok {
		l.evictedList.MoveToFront(item)
		return item.Value.(lruNode).val
	}
	return -1
}

func (l *LRU) Put(key, val int) {
	item, ok := l.items[key]
	if ok {
		item.Value = lruNode{
			key: key,
			val: val,
		}
		l.evictedList.MoveToFront(item)
	} else {
		if len(l.items) >= l.capacity {
			evicted := l.evictedList.Back()
			delete(l.items, evicted.Value.(lruNode).key)
			l.evictedList.Remove(evicted)
		}
		newItem := l.evictedList.PushFront(lruNode{key,val})
		l.items[key] = newItem
	}
}

