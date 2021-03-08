package main

type LFU struct {
	capacity int
	kv map[int]int
}

func NewLFU(capacity int) *LFU {
	return &LFU{capacity: capacity}
}


func (l *LFU) Get(key int) int  {
	item, ok := l.items[key]
	if ok {
		l.evictedList.MoveToFront(item)
		return item.Value.(entry).val
	}
	return -1
}

func (l *LFU) Put(key, val int) {
}