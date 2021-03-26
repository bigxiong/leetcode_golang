package main

import (
	"container/list"
)

/**
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present.
When the cache reached its capacity,
it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

*/

type entry struct {
	key int
	val int
}

type LRUCache struct {
	items     map[int]*list.Element
	evictList *list.List
	capacity  int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		items:     make(map[int]*list.Element),
		evictList: list.New(),
		capacity:  capacity,
	}
	return cache
}

func (l *LRUCache) Get(key int) int {
	if item, exists := l.items[key]; exists {
		// 活跃的移动到最前
		l.evictList.MoveToFront(item)
		return item.Value.(entry).val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	// update
	if item, ok := l.items[key]; ok {
		item.Value = entry{
			key: key,
			val: value,
		}
		l.evictList.MoveToFront(item)
	} else {
		if l.evictList.Len() >= l.capacity {
			tailItem := l.evictList.Back()
			delete(l.items, tailItem.Value.(entry).key)
			l.evictList.Remove(tailItem)
		}
		newItem := l.evictList.PushFront(entry{
			key: key,
			val: value,
		})
		l.items[key] = newItem
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
