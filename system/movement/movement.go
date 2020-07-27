package movement

import (
	"taylz.io/game/component/collision"
	"taylz.io/game/component/move"
	"taylz.io/game/component/shape"
	"taylz.io/game/entity"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type S struct {
	Collision *collision.Cache
	Shape     *shape.Cache
	Move      *move.Cache
}

// New returns a new movement system
func New(collision *collision.Cache, shape *shape.Cache, move *move.Cache) *S {
	return &S{
		Collision: collision,
		Shape:     shape,
		Move:      move,
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

		s.Move.Each(func(entity entity.T, move *move.T) {
			if move == nil {
			} else if move.I == nil {
			} else if shape := s.Shape.Get(entity); shape == nil {
			} else {
				v := move.Next(entity).MakeUnit().Multiply(d)
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
