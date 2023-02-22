package main

import (
	"github.com/rhcourses/go.structs/examples/cardgames/maumau"
)

func main() {
	g := maumau.New("Arno Nym", "Bye Speel")
	g.Setup()
	g.Run()
	g.PrintState()
}
