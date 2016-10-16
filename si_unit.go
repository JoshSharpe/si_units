package si_units

import "fmt"

type SIUnit struct {
	Id          string
	Name        string
	Value       float64
	VectorValue *Vector
	Unit        string
}

func (u *SIUnit) String() string {
	if u.VectorValue == nil {
		return fmt.Sprintf("%f %s", u.Value, u.Unit)
	}
	return fmt.Sprintf("%f %s", u.VectorValue, u.Unit)
}

func (u *SIUnit) Add(x float64) {
	u.Value = u.Value + x
}

func (u *SIUnit) Subtract(x float64) {
	u.Value = u.Value - x
}

func (u *SIUnit) Multiply(x float64) {
	u.Value = u.Value * x
}

func (u *SIUnit) Divide(x float64) {
	u.Value = u.Value / x
}

func (u *SIUnit) AddVec(vec *Vector) {
	u.VectorValue = u.VectorValue.Add(vec)
}

func (u *SIUnit) SubtractVec(vec *Vector) {
	u.VectorValue = u.VectorValue.Subtraction(vec)
}

func (u *SIUnit) MultiplyVec(scalar float64) {
	u.VectorValue = u.VectorValue.Multiply(scalar)
}

func (u *SIUnit) DivideVec(scalar float64) {
	u.VectorValue = u.VectorValue.Divide(scalar)
}

type Mass struct {
	SIUnit
}

func NewMass(name string, mass float64) *Mass {
	id := "mass-" + name
	return &Mass{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: mass,
			Unit:  "kg",
		},
	}
}

func (m *Mass) MultiplyAcceleration(a *Acceleration) *Force {
	return NewForce(
		m.Name+"*"+a.Name,
		a.VectorValue.Multiply(m.Value),
	)
}

type Length struct {
	SIUnit
}

func NewLength(name string, length float64) *Length {
	id := "length-" + name
	return &Length{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: length,
			Unit:  "m",
		},
	}
}

type Time struct {
	SIUnit
}

func NewTime(name string, time float64) *Time {
	id := "time-" + name
	return &Time{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: time,
			Unit:  "s",
		},
	}
}
