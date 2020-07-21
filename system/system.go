package system

import (
	"taylz.io/game/component"
	"taylz.io/game/system/movement"
)

type S struct {
	Movement *movement.S
}

func New(c *component.C) *S {
	return &S{
		Movement: movement.New(c),
	}
}
