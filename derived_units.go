package SIUnit

type Area SIUnit

func NewArea(name string, area float64) *Area {
	id := "area-" + name
	return &Area{
		Id:    id,
		Name:  name,
		Value: area,
		Unit:  "m^2",
	}
}

func CalculateNewArea(name string, num1 *Length, num2 *Length) *Area {
	id := "area-" + name
	return &Area{
		Id:    id,
		Name:  name,
		Value: num1.Value * num2.Value,
		Unit:  "m^2",
	}
}

type Force SIUnit

func NewForce(name string, f *Vector) *Force {
	id := "force-" + name
	return &Force{
		Id:          id,
		Name:        name,
		VectorValue: f,
		Unit:        "N",
	}
}

type Velocity SIUnit

func NewVelocity(name string, v *Vector) *Velocity {
	id := "velocity-" + name
	return &Velocity{
		Id:          id,
		Name:        name,
		VectorValue: v,
		Unit:        "m/s",
	}
}

type Acceleration SIUnit

func NewAcceleration(name string, v *Vector) *Acceleration {
	id := "acceleration-" + name
	return &Acceleration{
		Id:          id,
		Name:        name,
		VectorValue: v,
		Unit:        "m/s^2",
	}
}
