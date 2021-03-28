package main

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (lruCache *LRUCache) Get(key int) int {
	cache := lruCache.m
	if v, exist := cache[key]; exist {
		lruCache.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (lruCache *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (lruCache *LRUCache) AddNode(node *LinkNode) {
	head := lruCache.head
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (lruCache *LRUCache) MoveToHead(node *LinkNode) {
	lruCache.RemoveNode(node)
	lruCache.AddNode(node)
}

func (lruCache *LRUCache) Put(key int, value int) {
	tail := lruCache.tail
	cache := lruCache.m
	if v, exist := cache[key]; exist {
		v.val = value
		lruCache.MoveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == lruCache.cap {
			delete(cache, tail.pre.key)
			lruCache.RemoveNode(tail.pre)
		}
		lruCache.AddNode(v)
		cache[key] = v
	}
}
