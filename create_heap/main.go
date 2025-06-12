package main

import (
	"container/heap"
	"fmt"
)

type SmallData struct {
	name string
	size int
}
type Item struct {
	id   int
	data *SmallData
}
type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].id < h[j].id }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ItemHeap) Push(x any) {
	*h = append(*h, x.(*Item))
}
func (h *ItemHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func main() {
	h := &ItemHeap{}
	heap.Init(h)
	heap.Push(h, &Item{id: 5, data: &SmallData{name: "five", size: 50}})
	heap.Push(h, &Item{id: 2, data: &SmallData{name: "two", size: 20}})
	heap.Push(h, &Item{id: 8, data: &SmallData{name: "eight", size: 80}})
	heap.Push(h, &Item{id: 1, data: &SmallData{name: "one", size: 10}})
	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		fmt.Printf("id: %d, name: %s, size: %d\n", item.id, item.data.name, item.data.size)
	}
}
