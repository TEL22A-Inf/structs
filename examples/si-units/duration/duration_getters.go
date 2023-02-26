package duration

// ToMS liefert die Millisekunden dieser Duration als int.
func (d Duration) ToMS() int {
	return int(d)
}

// ToS liefert die Sekunden dieser Duration als int.
func (d Duration) ToS() int {
	return d.ToMS() / 1000
}

// ToM liefert die Minuten dieser Duration als int.
func (d Duration) ToM() int {
	return d.ToS() / 60
}

// ToH liefert die Stunden dieser Duration als int.
func (d Duration) ToH() int {
	return d.ToM() / 60
}

// ToD liefert die Stunden dieser Duration als int.
func (d Duration) ToD() int {
	return d.ToH() / 24
}
