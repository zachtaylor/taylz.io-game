package will

import (
	"taylz.io/game/entity"
	"taylz.io/types"
)

// C is a component
type C struct {
	cache map[entity.T]*T
	mu    types.Mutex
	obs   []F
}

// F is a func
type F = func(entity.T, *T)

// New returns a new Component
func New() *C {
	return &C{
		cache: make(map[entity.T]*T),
		obs:   make([]F, 0),
	}
}

// Get returns the component for an entity
func (c *C) Get(entity entity.T) *T { return c.cache[entity] }

// Set saves a component for an entity
func (c *C) Set(entity entity.T, t *T) {
	c.mu.Lock()
	if t != nil {
		c.cache[entity] = t
	} else {
		delete(c.cache, entity)
	}
	c.mu.Unlock()
	for _, f := range c.obs {
		f(entity, t)
	}
}

// Each calls the func for each entity with this component
func (c *C) Each(f F) {
	c.mu.Lock()
	for k, v := range c.cache {
		f(k, v)
	}
	c.mu.Unlock()
}

// Keys returns a new slice with all the entity ids with component data
func (c *C) Keys() []entity.T {
	c.mu.Lock()
	keys := make([]entity.T, 0, len(c.cache))
	for k := range c.cache {
		keys = append(keys, k)
	}
	c.mu.Unlock()
	return keys
}

// Observe adds a func to be called when a component changes
func (c *C) Observe(f F) {
	c.obs = append(c.obs, f)
}

// Remove deletes an entity component
func (c *C) Remove(entity entity.T) { c.Set(entity, nil) }
