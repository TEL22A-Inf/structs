package measures

import "fmt"

// ExampleVelocity_constructors konstruiert verschiedene Velocity-Objekte und gibt sie aus.
func ExampleVelocity_constructors() {
	// Geschwindigkeit aus km/h-Angabe:
	v1 := KMH(42)
	fmt.Println(v1)

	// Geschwindigkeit aus m/s-Angabe:
	v2 := MS(15)
	fmt.Println(v2)

	// Geschwindigkeit aus Meilen pro 2 Tage:
	distance := Miles(50)
	duration := Days(2)
	v3 := NewVelocity(distance, duration)
	fmt.Println(v3)

	// Output:
	// {{42000000} {3600000}}
	// {{15000} {1000}}
	// {{80467000} {172800000}}
}
