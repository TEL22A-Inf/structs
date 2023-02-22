package measures

import "fmt"

// ExampleDuration_constructors erzeugt fünf verschiedene Durations jeweils mit
// einem der verschiedenen Konstruktoren.
// Jedes erzeugte Objekt wird unmittelbar danach ausgegeben.
func ExampleDuration_constructors() {
	// Erzeuge jeweils ein Duration-Objekt mit einem der Konstruktoren und gebe es aus:

	d1 := Milliseconds(1)
	fmt.Println(d1)

	d2 := Seconds(1)
	fmt.Println(d2)

	d3 := Minutes(1)
	fmt.Println(d3)

	d4 := Hours(1)
	fmt.Println(d4)

	d5 := Days(1)
	fmt.Println(d5)

	// In den Ausgaben unten sehen wir, dass das Struct
	// immer eine Millisekunden-Anzahl enthält.

	// Output:
	// {1}
	// {1000}
	// {60000}
	// {3600000}
	// {86400000}
}
