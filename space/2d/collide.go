package space

// CollideRR returns whether 2 rectangles collide
func CollideRR(r1, r2 Rect) bool {
	colx := (r1.Point.X < r2.Point.X && r1.Point.X+r1.Size.DX >= r2.Point.X) ||
		(r1.Point.X > r2.Point.X && r1.Point.X <= r2.Point.X+r2.Size.DX)
	coly := (r1.Point.Y < r2.Point.Y && r1.Point.Y+r1.Size.DY >= r2.Point.Y) ||
		(r1.Point.Y > r2.Point.Y && r1.Point.Y <= r2.Point.Y+r2.Size.DY)
	return colx && coly
}
