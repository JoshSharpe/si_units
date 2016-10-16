package si_units

type Area struct {
	SIUnit
}

func NewArea(name string, area float64) *Area {
	id := "area-" + name
	return &Area{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: area,
			Unit:  "m^2",
		},
	}
}

func CalculateNewArea(name string, num1 *Length, num2 *Length) *Area {
	id := "area-" + name
	return &Area{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: num1.Value * num2.Value,
			Unit:  "m^2",
		},
	}
}

type Volume struct {
	SIUnit
}

func NewVolume(name string, vol float64) *Volume {
	id := "volume-" + name
	return &Volume{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: vol,
			Unit:  "m^3",
		},
	}
}

type Density struct {
	SIUnit
}

func NewDensity(name string, density float64) *Density {
	id := "density-" + name
	return &Density{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: density,
			Unit:  "kg/m^3",
		},
	}
}

func (d *Density) MultiplyVolume(v *Volume) *Mass {
	return NewMass(
		d.Name+"*"+v.Name,
		d.Value*v.Value,
	)
}

type FlowRate struct {
	SIUnit
}

func NewFlowRate(name string, flowRate float64) *FlowRate {
	id := "flowRate-" + name
	return &FlowRate{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: flowRate,
			Unit:  "m^3/s",
		},
	}
}

func (f *FlowRate) MultiplyTime(t *Time) *Volume {
	return NewVolume(
		f.Name+"*"+t.Name,
		f.Value*t.Value,
	)
}

type MolarMass struct {
	SIUnit
}

func NewMolarMass(name string, molarMass float64) *MolarMass {
	id := "molarMass-" + name
	return &MolarMass{
		SIUnit{
			Id:    id,
			Name:  name,
			Value: molarMass,
			Unit:  "kg/mol",
		},
	}
}

// Vector units
type Force struct {
	SIUnit
}

func NewForce(name string, f *Vector) *Force {
	id := "force-" + name
	return &Force{
		SIUnit{
			Id:          id,
			Name:        name,
			VectorValue: f,
			Unit:        "N",
		},
	}
}

type Displacement struct {
	SIUnit
}

func NewDisplacement(name string, v *Vector) *Displacement {
	id := "displacement-" + name
	return &Displacement{
		SIUnit{
			Id:          id,
			Name:        name,
			VectorValue: v,
			Unit:        "m",
		},
	}
}

type Velocity struct {
	SIUnit
}

func NewVelocity(name string, v *Vector) *Velocity {
	id := "velocity-" + name
	return &Velocity{
		SIUnit{
			Id:          id,
			Name:        name,
			VectorValue: v,
			Unit:        "m/s",
		},
	}
}

func (v *Velocity) MultiplyTime(t *Time) *Displacement {
	return NewDisplacement(
		v.Name+"*"+t.Name,
		v.VectorValue.Multiply(t.Value),
	)
}

type Acceleration struct {
	SIUnit
}

func NewAcceleration(name string, v *Vector) *Acceleration {
	id := "acceleration-" + name
	return &Acceleration{
		SIUnit{
			Id:          id,
			Name:        name,
			VectorValue: v,
			Unit:        "m/s^2",
		},
	}
}

func (a *Acceleration) MultiplyTime(t *Time) *Velocity {
	return NewVelocity(
		a.Name+"*"+t.Name,
		a.VectorValue.Multiply(t.Value),
	)
}
