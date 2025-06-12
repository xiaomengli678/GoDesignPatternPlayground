package main

import "fmt"

type evictionBase interface {
	evict(queue *[]*Node)
}

type evictFromHead struct{}

func (h *evictFromHead) evict(queue *[]*Node) {
	if len(*queue) > 0 {
		*queue = (*queue)[1:]
	}
}

type evictFromTail struct{}

func (t *evictFromTail) evict(queue *[]*Node) {
	if len(*queue) > 0 {
		*queue = (*queue)[:len(*queue)-1]
	}
}

type Node struct {
	key string
	val string
}

type cacheSystem struct {
	queue         []*Node
	evictStrategy evictionBase
}

func (c *cacheSystem) set(key string, value string) {
	c.queue = append(c.queue, &Node{key: key, val: value})
}
func (c *cacheSystem) get(key string) (string, bool) {
	for _, n := range c.queue {
		if n.key == key {
			return n.val, true
		}
	}
	return "", false
}
func (c *cacheSystem) evict() {
	c.evictStrategy.evict(&c.queue)
}

func main() {
	cs0 := cacheSystem{queue: []*Node{}, evictStrategy: &evictFromHead{}}
	cs0.set("hello", "world")
	cs0.set("helloagain", "worldagain")
	cs0.evict()
	for _, n := range cs0.queue {
		fmt.Println(n.key, n.val)
	}
	fmt.Println("----------------------")
	cs1 := cacheSystem{queue: []*Node{}, evictStrategy: &evictFromTail{}}
	cs1.set("hello", "world")
	cs1.set("helloagain", "worldagain")
	cs1.evict()
	for _, n := range cs1.queue {
		fmt.Println(n.key, n.val)
	}
}
