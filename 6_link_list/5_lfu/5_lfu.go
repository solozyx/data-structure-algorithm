package main

type LFUCache struct {
	cache               map[int]*Node
	freq                map[int]*DoubleList
	ncap, size, minFreq int
}

func (lfuCache *LFUCache) IncrFreq(node *Node) {
	_freq := node.freq
	lfuCache.freq[_freq].RemoveNode(node)
	if lfuCache.minFreq == _freq && lfuCache.freq[_freq].IsEmpty() {
		lfuCache.minFreq++
		delete(lfuCache.freq, _freq)
	}
	node.freq++

	if lfuCache.freq[node.freq] == nil {
		lfuCache.freq[node.freq] = createDL()
	}
	lfuCache.freq[node.freq].AddFirst(node)
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache: make(map[int]*Node),
		freq:  make(map[int]*DoubleList),
		ncap:  capacity,
	}
}

func (lfuCache *LFUCache) Get(key int) int {
	if node, ok := lfuCache.cache[key]; ok {
		lfuCache.IncrFreq(node)
		return node.val
	}

	return -1
}

func (lfuCache *LFUCache) Put(key int, value int) {
	if lfuCache.ncap == 0 {
		return
	}
	//节点存在
	if node, ok := lfuCache.cache[key]; ok {
		node.val = value
		lfuCache.IncrFreq(node)
	} else {
		if lfuCache.size >= lfuCache.ncap {
			node := lfuCache.freq[lfuCache.minFreq].RemoveLast()
			delete(lfuCache.cache, node.key)
			lfuCache.size--
		}
		x := &Node{key: key, val: value, freq: 1}
		lfuCache.cache[key] = x
		if lfuCache.freq[1] == nil {
			lfuCache.freq[1] = createDL()
		}
		lfuCache.freq[1].AddFirst(x)
		lfuCache.minFreq = 1
		lfuCache.size++
	}
}

//节点node
type Node struct {
	key, val, freq int
	prev, next     *Node
}

//双链表
type DoubleList struct {
	tail, head *Node
}

//创建一个双链表
func createDL() *DoubleList {
	head, tail := &Node{}, &Node{}
	head.next, tail.prev = tail, head

	return &DoubleList{
		tail: tail,
		head: head,
	}
}

func (lfuCache *DoubleList) IsEmpty() bool {
	return lfuCache.head.next == lfuCache.tail
}

//将node添加为双链表的第一个元素
func (lfuCache *DoubleList) AddFirst(node *Node) {
	node.next = lfuCache.head.next
	node.prev = lfuCache.head

	lfuCache.head.next.prev = node
	lfuCache.head.next = node
}

func (lfuCache *DoubleList) RemoveNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next

	node.next = nil
	node.prev = nil
}

func (lfuCache *DoubleList) RemoveLast() *Node {
	if lfuCache.IsEmpty() {
		return nil
	}

	lastNode := lfuCache.tail.prev
	lfuCache.RemoveNode(lastNode)

	return lastNode
}
