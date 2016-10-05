package SIUnit

import "fmt"

type SIUnit struct {
	Id          string
	Name        string
	Value       float64
	VectorValue *Vector
	Unit        string
}

func (u *SIUnit) String() string {
	return fmt.Sprintf("%f %s", u.Value, u.Unit)
}

type Mass SIUnit

func NewMass(name string, mass float64) *Mass {
	id := "mass-" + name
	return &Mass{
		Id:    id,
		Name:  name,
		Value: mass,
		Unit:  "kg",
	}
}

type Length SIUnit

func NewLength(name string, length float64) *Length {
	id := "length-" + name
	return &Length{
		Id:    id,
		Name:  name,
		Value: length,
		Unit:  "m",
	}
}
