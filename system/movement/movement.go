package movement

import (
	"taylz.io/game/component/collision"
	"taylz.io/game/component/shape"
	"taylz.io/game/component/will"
	"taylz.io/game/entity"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type S struct {
	Collision *collision.C
	Shape     *shape.C
	Will      *will.C
}

// New returns a new movement system
func New(collision *collision.C, shape *shape.C, will *will.C) *S {
	return &S{
		Collision: collision,
		Shape:     shape,
		Will:      will,
	}
}

func (s *S) Run(f types.Duration) {
	go s.run(f)
}

func (s *S) run(f types.Duration) {
	d := float64(f) / float64(types.Second)
	// e := types.NewTime()
	for range types.NewChanTick(f) {
		// t := types.NewTime()
		// d := float64(t.Sub(e)) / float64(types.Second)
		// e = t

		s.Will.Each(func(entity entity.T, will *will.T) {
			if will == nil {
			} else if will.Move() == nil {
			} else if shape := s.Shape.Get(entity); shape == nil {
			} else {
				v := will.Move().Next(entity).MakeUnit().Multiply(d)
				p := shape.Point.Move(v)

				if s.collides(entity, space.Rect{
					Point: p,
					Size:  shape.Size,
				}) {
					return
				}

				shape.Point = p
			}
		})
	}
}

func (s *S) collides(entity entity.T, r space.Rect) bool {
	for _, k := range s.Collision.Keys() {
		if k == entity {
		} else if shape := s.Shape.Get(k); shape == nil {
		} else if space.CollideRR(r, shape.Rect) {
			return true
		}
	}

	return false
}
