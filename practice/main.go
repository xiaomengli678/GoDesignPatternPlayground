package main

import "fmt"
import "sync"

const capacity = 2

var once sync.Once
var instance Cache

func MakeMeFIFOCache() Cache {
	once.Do(func() {
		instance = &FIFOCache{Record: map[string]string{}, evictStrategy: &EvictHeadStrategy{}}
	})
	return instance
}

type Cache interface {
	set(key string, val string) bool
	get(key string) (string, bool)
}

type EvictStrategy interface {
	Evict(f *FIFOCache)
}

type FIFOCache struct {
	Record        map[string]string
	Queue         []string
	evictStrategy EvictStrategy
}

func (f *FIFOCache) set(key string, val string) bool {
	if _, ok := f.Record[key]; !ok {
		if len(f.Queue) == capacity {
			f.doSomeEviction()
			f.evictStrategy = &EvictTailStrategy{}
		}
		f.Record[key] = val
		f.Queue = append(f.Queue, key)
		return true
	}
	f.Record[key] = val
	return false
}
func (f *FIFOCache) get(key string) (string, bool) {
	if _, ok := f.Record[key]; !ok {
		return "", false
	}
	return f.Record[key], true
}
func (f *FIFOCache) doSomeEviction() {
	f.evictStrategy.Evict(f)
}

type EvictHeadStrategy struct{}

func (e *EvictHeadStrategy) Evict(f *FIFOCache) {
	key := f.Queue[0]
	f.Queue = f.Queue[1:]
	delete(f.Record, key)
}

type EvictTailStrategy struct{}

func (e *EvictTailStrategy) Evict(f *FIFOCache) {
	key := f.Queue[len(f.Queue)-1]
	f.Queue = f.Queue[:len(f.Queue)-1]
	delete(f.Record, key)
}

func main() {
	f := MakeMeFIFOCache()
	f.set("a", "1")
	f.set("b", "2")
	f.set("c", "3") // should evict "a" if using FIFO
	v, ok := f.get("a")
	fmt.Println("a:", v, ok) // expect "", false
}
