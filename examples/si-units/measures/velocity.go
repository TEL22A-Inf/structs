package measures

// Velocity ist ein Typ, der Geschwindigkeiten speichert.
// Eine Geschwindigkeit ist ein Quotient aus Entfernung und Zeit,
// daher werden hier Distance und Duration verwendet.
type Velocity struct {
	distance Distance
	duration Duration
}

// NewVelocity erwartet eine Distance und eine Duration und
// konstruiert daraus ein Velocity-Objekt.
func NewVelocity(distance Distance, duration Duration) Velocity {
	return Velocity{distance: distance, duration: duration}
}

// KMH liefert ein Velocity-Objekt zur gegebenen KM-Pro-Stunde-Angabe.
func KMH(kmh int) Velocity {
	return NewVelocity(Kilometers(kmh), Hours(1))
}

// MPS liefert ein Velocity-Objekt zur gegebenen Meter-Pro-Sekunde-Angabe.
func MS(ms int) Velocity {
	return NewVelocity(Meters(ms), Seconds(1))
}
