package main

import "fmt"

const capacity = 2

// Strategy interface: encapsulates the eviction logic
type EvictionStrategy interface {
	Evict(c *QueueCache)
}

// Node represents a key-value entry
type Node struct {
	key int
	val string
}

// Cache interface (optional for abstraction/testing)
type GeneralCache interface {
	SetKey(key int, val string) bool
	GetKey(key int) string
}

// QueueCache is the main cache system
type QueueCache struct {
	queue         []int            // order of insertion
	record        map[int]*Node    // actual key-value store
	evictStrategy EvictionStrategy // policy
}

// SetKey inserts or updates a key
func (c *QueueCache) SetKey(key int, val string) bool {
	if node, ok := c.record[key]; ok {
		node.val = val
		return true
	}
	if len(c.queue) == capacity {
		c.evictStrategy.Evict(c)
	}
	c.record[key] = &Node{key: key, val: val}
	c.queue = append(c.queue, key)
	return true
}

// GetKey retrieves a value by key
func (c *QueueCache) GetKey(key int) string {
	if node, ok := c.record[key]; ok {
		return node.val
	}
	return ""
}

// Eviction Strategies

type EvictHead struct{}

func (eh *EvictHead) Evict(c *QueueCache) {
	head := c.queue[0]
	c.queue = c.queue[1:]
	delete(c.record, head)
	fmt.Printf("Evicted (head): %d\n", head)
}

type EvictTail struct{}

func (et *EvictTail) Evict(c *QueueCache) {
	tail := c.queue[len(c.queue)-1]
	c.queue = c.queue[:len(c.queue)-1]
	delete(c.record, tail)
	fmt.Printf("Evicted (tail): %d\n", tail)
}

// --- Main ---
func main() {
	fmt.Println("EvictHead policy:")
	cacheHead := QueueCache{
		queue:         []int{},
		record:        make(map[int]*Node),
		evictStrategy: &EvictHead{},
	}
	cacheHead.SetKey(1, "hello")
	cacheHead.SetKey(2, "good bye")
	cacheHead.SetKey(3, "good day")
	for k, n := range cacheHead.record {
		fmt.Println(k, ":", n.val)
	}

	fmt.Println("\nEvictTail policy:")
	cacheTail := QueueCache{
		queue:         []int{},
		record:        make(map[int]*Node),
		evictStrategy: &EvictTail{},
	}
	cacheTail.SetKey(1, "hello")
	cacheTail.SetKey(2, "good bye")
	cacheTail.SetKey(3, "good day")
	for k, n := range cacheTail.record {
		fmt.Println(k, ":", n.val)
	}
}
