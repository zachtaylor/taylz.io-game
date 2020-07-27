package mover

import (
	"taylz.io/game/component/move"
	"taylz.io/game/component/shape"
	"taylz.io/game/entity"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type Follow struct {
	Shape  *shape.C
	Entity entity.T
}

func (f *Follow) isMover() move.I { return f }

func (f *Follow) Next(entity entity.T) space.Vector {
	src := f.Shape.Get(entity)
	dest := f.Shape.Get(f.Entity)
	return space.Vector{
		DX: dest.Point.X - src.Point.X,
		DY: dest.Point.Y - src.Point.Y,
	}
}

func (f *Follow) Data() types.Dict {
	return types.Dict{
		"f": types.Dict{
			"e": f.Entity,
		},
	}
}
