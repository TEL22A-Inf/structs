package velocity

// ToKMH liefert die Stundenkilometer als int.
func (v Velocity) ToKMH() int {
	return (v.ToMMpMSS() * (1000 * 60 * 60)) / (1000 * 1000)
}

// ToMPS liefert die Meter pro Sekunde als int.
func (v Velocity) ToMPS() int {
	return v.ToMMpMSS()
}

// ToMMpMSS liefert die Millimeter pro Millisekunde als int.
func (v Velocity) ToMMpMSS() int {
	return v.dist.ToMM() / v.dur.ToMS()
}
