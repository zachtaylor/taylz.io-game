package space

type Point struct {
	X float64
	Y float64
}

func (p *Point) Copy() Point {
	return Point{
		X: p.X,
		Y: p.Y,
	}
}

func (p *Point) Move(v Vector) Point {
	return Point{
		X: p.X + v.DX,
		Y: p.Y + v.DY,
	}
}
