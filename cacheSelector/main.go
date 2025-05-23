package main

type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

type BasicCache interface {
	Get(key int) int
	Put(key int, value int)
}

type LRUCache struct {
	head   *Node
	tail   *Node
	cap    int
	record map[int]*Node
}

func LRUConstructor(capacity int) BasicCache {
	head := &Node{key: 0, val: 0}
	tail := &Node{key: 0, val: 0}
	head.next = tail
	tail.prev = head
	record := make(map[int]*Node)
	return &LRUCache{
		head:   head,
		tail:   tail,
		cap:    capacity,
		record: record}
}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.record[key]; !ok {
		return -1
	}
	oldNode := l.record[key]
	newNode := &Node{key: key, val: oldNode.val}
	l.removeNode(oldNode)
	l.addNode(newNode)
	delete(l.record, key)
	l.record[key] = newNode
	return newNode.val
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.record[key]; ok {
		oldNode := l.record[key]
		newNode := &Node{key: key, val: value}
		l.removeNode(oldNode)
		l.addNode(newNode)
		delete(l.record, key)
		l.record[key] = newNode
	} else {
		if len(l.record) == l.cap {
			oldNode := l.tail.prev
			l.removeNode(oldNode)
			delete(l.record, oldNode.key)
		}
		newNode := &Node{key: key, val: value}
		l.addNode(newNode)
		l.record[key] = newNode
	}
}

func (l *LRUCache) removeNode(node *Node) {
	prev_ := node.prev
	next_ := node.next
	prev_.next = next_
	next_.prev = prev_
}

func (l *LRUCache) addNode(node *Node) {
	oldNext := l.head.next
	l.head.next = node
	node.next = oldNext
	oldNext.prev = node
	node.prev = l.head
}

type FIFOCache struct {
	cap    int
	record []*Node
}

func FIFOConstructor(capacity int) BasicCache {
	record := []*Node{}
	return &FIFOCache{
		cap:    capacity,
		record: record}
}

func (f *FIFOCache) Get(key int) int {
	for _, node := range f.record {
		if node.key == key {
			return node.val
		}
	}
	return -1
}

func (f *FIFOCache) Put(key int, value int) {
	for _, node := range f.record {
		if node.key == key {
			node.val = value
			return
		}
	}
	new_ := &Node{key: key, val: value}
	if len(f.record) < f.cap {
		f.record = append([]*Node{new_}, f.record...)
	} else {
		f.record = f.record[1:]
		f.record = append(f.record, new_)
	}
}

func findCache(s string) BasicCache {
	switch s {
	case "lru":
		return LRUConstructor(10)
	case "fifo":
		return FIFOConstructor(10)
	default:
		return LRUConstructor(10)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
