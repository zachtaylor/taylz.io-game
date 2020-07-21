package component

import (
	"taylz.io/game/component/collision"
	"taylz.io/game/component/shape"
	"taylz.io/game/component/update"
	"taylz.io/game/component/will"
)

type C struct {
	Collision *collision.C
	Shape     *shape.C
	Update    *update.C
	Will      *will.C
}

func New() *C {
	return &C{
		Collision: collision.New(),
		Shape:     shape.New(),
		Update:    update.New(),
		Will:      will.New(),
	}
}
