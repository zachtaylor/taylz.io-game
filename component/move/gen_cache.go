// Code generated by go-gengen(v0.0.1) DO NOT EDIT.

package move

import (
	"sync"

	"taylz.io/game/entity"
)

// Cache is an observable concurrent in-memory datastore
type Cache struct {
	dat   map[entity.T]*T
	mu    sync.Mutex
	obs   []Func
}

// Storer is an abstraction of in-memory datastore
type Storer interface {
	Get(entity.T) *T
	Set(entity.T, *T)
	Each(Func)
	Sync(func(map[entity.T]*T))
	Keys() []entity.T
	Observe(Func)
	Remove(entity.T)
}

// Func is a callback func
type Func = func(entity.T, *T)

// NewCache returns a new Cache
func NewCache() *Cache {
	return &Cache{
		dat: make(map[entity.T]*T),
		obs:   make([]Func, 0),
	}
}

// Get returns the *T for a entity.T
func (c *Cache) Get(k entity.T) *T { return c.dat[k] }

// Set saves a *T for a entity.T
func (c *Cache) Set(k entity.T, v *T) {
	c.mu.Lock()
	if v != nil {
		c.dat[k] = v
	} else {
		delete(c.dat, k)
	}
	c.mu.Unlock()
	for _, f := range c.obs {
		f(k, v)
	}
}

// Each calls the func for each entity.T,*T in this Cache
func (c *Cache) Each(f Func) {
	c.mu.Lock()
	for k, v := range c.dat {
		f(k, v)
	}
	c.mu.Unlock()
}

// Sync calls the func within the cache lock state
func (c *Cache) Sync(f func(map[entity.T]*T)) {
	c.mu.Lock()
	f(c.dat)
	c.mu.Unlock()
}

// Keys returns a new slice with all the entity.T keys
func (c *Cache) Keys() []entity.T {
	c.mu.Lock()
	keys := make([]entity.T, 0, len(c.dat))
	for k := range c.dat {
		keys = append(keys, k)
	}
	c.mu.Unlock()
	return keys
}

// Observe adds a func to be called when a *T is explicitly set
func (c *Cache) Observe(f Func) { c.obs = append(c.obs, f) }

// Remove deletes a entity.T,*T
func (c *Cache) Remove(k entity.T) { c.Set(k, nil) }
