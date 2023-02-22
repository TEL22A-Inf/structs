package measures

// Distance ist ein Datentyp für Entfernungen.
// Als Basis-Einheit wird eine Entfernung in Millimetern gespeichert.
type Distance struct {
	mm int
}

// Millimeters konstruiert eine Entfernung aus einer Millimeter-Angabe.
func Millimeters(mm int) Distance {
	return Distance{mm}
}

// Centimeters konstruiert eine Entfernung aus einer Zentimeter-Angabe.
func Centimeters(cm int) Distance {
	return Millimeters(cm * 10)
}

// Meters konstruiert eine Entfernung aus einer Meter-Angabe.
func Meters(m int) Distance {
	return Centimeters(m * 100)
}

// Kilometers konstruiert eine Entfernung aus einer Kilometer-Angabe.
func Kilometers(km int) Distance {
	return Meters(km * 1000)
}

// Kilometers konstruiert eine Entfernung aus einer Meilen-Angabe.
func Miles(mi int) Distance {
	return Centimeters(mi * 160934)
}
