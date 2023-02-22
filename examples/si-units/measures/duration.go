package measures

// Duration ist ein Datentyp für Zeitdauern.
// Als Basis-Einheit wird eine Zeitdauer in Millisekunden gespeichert.
type Duration struct {
	ms int
}

// Milliseconds konstruiert eine Duration aus einer Millisekunden-Angabe.
func Milliseconds(ms int) Duration {
	return Duration{ms}
}

// Seconds konstruiert eine Duration aus einer Sekunden-Angabe.
func Seconds(s int) Duration {
	return Milliseconds(s * 1000)
}

// Minutes konstruiert eine Duration aus einer Minuten-Angabe.
func Minutes(m int) Duration {
	return Seconds(m * 60)
}

// Hours konstruiert eine Duration aus einer Stunden-Angabe.
func Hours(h int) Duration {
	return Minutes(h * 60)
}

// Days konstruiert eine Duration aus einer Tages-Angabe.
func Days(d int) Duration {
	return Hours(d * 24)
}
