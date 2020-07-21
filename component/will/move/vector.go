package move

import (
	"taylz.io/game/world/space"
	"taylz.io/types"
)

type Vector space.Vector

func (v Vector) Next() space.Vector { return space.Vector(v) }
func (v Vector) Data() types.Dict {
	return types.Dict{
		"v": types.Dict{
			"x": v.DX,
			"y": v.DY,
		},
	}
}

// VectorNorth is the unit vector up
var VectorNorth = Vector{0, 1}

// VectorEast is the unit vector right
var VectorEast = Vector{1, 0}

// VectorSouth is the unit vector down
var VectorSouth = Vector{0, -1}

// VectorWest is the unit vector left
var VectorWest = Vector{-1, 0}
