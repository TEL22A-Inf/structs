package distance

// ToMM liefert die Millimeter dieser Distance als int.
func (d Distance) ToMM() int {
	return int(d)
}

// ToCM liefert die Zentimeter dieser Distance als int.
func (d Distance) ToCM() int {
	return d.ToMM() / 10
}

// ToM liefert die Meter dieser Distance als int.
func (d Distance) ToM() int {
	return d.ToCM() / 100
}

// ToKM liefert die Kilometer dieser Distance als int.
func (d Distance) ToKM() int {
	return d.ToM() / 1000
}

// ToMi liefert die Meilen dieser Distance als int.
func (d Distance) ToMi() int {
	return d.ToCM() / 160934
}
