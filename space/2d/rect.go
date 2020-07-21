package space

type Rect struct {
	Point Point
	Size  Vector
}

func (r *Rect) Copy() Rect {
	return Rect{
		Point: Point{
			X: r.Point.X,
			Y: r.Point.Y,
		},
		Size: Vector{
			r.Size.DX,
			r.Size.DY,
		},
	}
}
