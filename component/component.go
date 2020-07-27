package component

import (
	"taylz.io/game/component/collision"
	"taylz.io/game/component/move"
	"taylz.io/game/component/shape"
	"taylz.io/game/component/update"
)

type Collision = collision.Cache

var NewCollision = collision.NewCache

type Shape = shape.Cache

var NewShape = shape.NewCache

type Update = update.Cache

var NewUpdate = update.NewCache

type Move = move.Cache

var NewMove = move.NewCache
