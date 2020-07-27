package mover

import (
	"taylz.io/game/component/shape"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type Point struct {
	Src   *shape.T
	Point *space.Point
}

func (p *Point) Next() space.Vector {
	return space.Vector{
		DX: p.Point.X - p.Src.Point.X,
		DY: p.Point.Y - p.Src.Point.Y,
	}
}

func (p *Point) Data() types.Dict {
	return types.Dict{
		"p": types.Dict{
			"x": p.Point.X,
			"y": p.Point.Y,
		},
	}
}
