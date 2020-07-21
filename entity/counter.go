package entity

import "taylz.io/types"

type Counter struct {
	id T
	mu types.Mutex
}

func (c *Counter) Next() (entity T) {
	c.mu.Lock()
	c.id++
	entity = c.id
	c.mu.Unlock()
	return
}
