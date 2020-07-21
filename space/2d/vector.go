package space

import "math"

type Vector struct {
	DX float64
	DY float64
}

func (v Vector) Add(w Vector) Vector {
	return Vector{
		DX: v.DX + w.DX,
		DY: v.DY + w.DY,
	}
}

func (v Vector) Copy() Vector {
	return Vector{
		DX: v.DX,
		DY: v.DY,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt((v.DX * v.DX) + (v.DY * v.DY))
}

func (v Vector) Multiply(f float64) Vector {
	return Vector{
		DY: v.DY * f,
		DX: v.DX * f,
	}
}

func (v Vector) Reverse() Vector {
	return Vector{
		DY: -v.DY,
		DX: -v.DX,
	}
}

func (v Vector) MakeUnit() Vector {
	length := v.Length()

	if math.IsNaN(length) || length == 0 {
		return v.Copy()
	}

	return Vector{
		DY: v.DY / length,
		DX: v.DX / length,
	}
}

func (v Vector) AsSlope() float64 {
	if v.DX == 0 {
		if v.DY > 0 {
			return math.Inf(1)
		} else {
			return math.Inf(-1)
		}
	} else {
		return v.DY / v.DX
	}
}

func (v Vector) MakeRightAngle() Vector {
	return Vector{
		DY: v.DX,
		DX: -v.DY,
	}
}
