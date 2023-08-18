package main

import "fmt"

// EvictionAlgo Strategy interface
type EvictionAlgo interface {
	evict(c *Cache)
}

// Fifo Strategy

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting using fifo strategy")
}

// Lru Strategy

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

// Lfu Strategy

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
