package mathies

import "math"

type Vector2 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
    v1.X += v2.X
    v1.Y += v2.Y

    return v1
}

func (v1 Vector2) Sub(v2 Vector2) Vector2  {
	v1.X -= v2.X
	v1.Y -= v2.Y

	return v1
}

func (v1 Vector2) Mul(s float64) Vector2 {
	return Vector2{
		v1.X * s,
		v1.Y * s,
	}
}

func (v1 Vector2) Dot(v2 Vector2) float64 {
    return v1.X * v2.X + v1.Y * v2.Y
}

func (v1 Vector2) MagSqr() float64 {
	return v1.Dot(v1)
}

func (v1 Vector2) Mag() float64 {
	return math.Sqrt(v1.MagSqr())
}

func (v1 Vector2) Normalized() Vector2 {
	mag := v1.Mag()

	// TODO: make sure this is the right call to prevent nan
	if mag == 0 {
		return v1
	}

	return v1.Mul(1.0 / mag)
}

func (v1 Vector2) Project(v2 Vector2) float64 {
	return v1.Dot(v2.Normalized())
}