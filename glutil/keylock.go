package glutil

import (
	"sync"
)

var Lock glock = glock{selfLock: &sync.RWMutex{},
	lockMap:   make(map[string]*sync.Mutex),
	rwLockMap: make(map[string]*sync.RWMutex),
}

type glock struct {
	rwLockMap map[string]*sync.RWMutex
	lockMap   map[string]*sync.Mutex
	selfLock  *sync.RWMutex
}

func (g *glock) Get(key string) *sync.Mutex {
	g.selfLock.RLock()
	l, exists := g.lockMap[key]
	g.selfLock.RUnlock()
	if exists {
		return l
	}

	return g.createNew(key)
}

func (g *glock) createNew(key string) *sync.Mutex {
	g.selfLock.Lock()
	defer g.selfLock.Unlock()
	g.lockMap[key] = &sync.Mutex{}

	return g.lockMap[key]
}
