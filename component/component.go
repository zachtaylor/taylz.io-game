package component

import (
	"taylz.io/game/component/collision"
	"taylz.io/game/component/shape"
	"taylz.io/game/component/update"
	"taylz.io/game/component/will"
)

type Collision = collision.C

var NewCollision = collision.New

type Shape = shape.C

var NewShape = shape.New

type Update = update.C

var NewUpdate = update.New

type Will = will.C

var NewWill = will.New
