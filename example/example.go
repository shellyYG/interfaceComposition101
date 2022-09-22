package example
import (
	"fmt"
)

type Cache struct {
	storage map[string]string
	evictionAlgo EvictionAlgo
	capacity int
	maxCapacity int
}

type EvictionAlgo interface {
    evict(c *Cache)
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache {
		storage: storage,
		evictionAlgo: e,
		capacity: 0,
		maxCapacity: 2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}


func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type Fifo struct {

}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evict by fifo")
}

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
    fmt.Println("Evict by lru")
}

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
    fmt.Println("Evict by lfu")
}


func Run() {
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