package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	store map[string]cacheEntry
	mu	  *sync.Mutex
}

type CacheEntry struct {
	createdAt 	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache {
	
}

func (c Cache)Add(key string, val []byte) {
	c.store[key] = val
}