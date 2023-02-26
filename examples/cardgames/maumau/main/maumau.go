package main

import (
	"github.com/tel22a-inf/structs/examples/cardgames/maumau"
)

func main() {
	g := maumau.New("Arno Nym", "Bye Speel")
	g.Setup()
	g.Run()
	g.PrintState()
}
