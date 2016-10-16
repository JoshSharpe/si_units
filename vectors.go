package si_units

import (
	"fmt"
	"math"
)

// TODO Make X, Y, Z units
type Vector struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%f, %f, %f>", v.X, v.Y, v.Z)
}

func NewVector(x float64, y float64, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return NewVector(
		v.X+v2.X,
		v.Y+v2.Y,
		v.Z+v2.Z,
	)
}

func (v *Vector) Subtraction(v2 *Vector) *Vector {
	return NewVector(
		v.X-v2.X,
		v.Y-v2.Y,
		v.Z-v2.Z,
	)
}

func (v *Vector) Multiply(scalar float64) *Vector {
	return NewVector(
		v.X*scalar,
		v.Y*scalar,
		v.Z*scalar,
	)
}

func (v *Vector) Divide(scalar float64) *Vector {
	return NewVector(
		v.X/scalar,
		v.Y/scalar,
		v.Z/scalar,
	)
}

func (v *Vector) DotProduct(v2 *Vector) float64 {
	return (v.X * v2.X) +
		(v.Y * v2.Y) +
		(v.Z * v2.Z)
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(
		(v.X * v.X) +
			(v.Y * v.Y) +
			(v.Z * v.Z))
}

func (v *Vector) IsUnit() bool {
	mag := v.Magnitude()
	return 0.99 <= mag && mag <= 1.01
}
