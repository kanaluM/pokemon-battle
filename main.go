package main

import (
	// "fmt"
)

func main() {
	cynthia := []*Pokemon{NewPokemon("Spiritomb", false)}
	player := []*Pokemon{NewPokemon("Torterra", false)}

	userOneInput := &UserInput{"attack", cynthia[0], cynthia, ""}
	userTwoInput := &UserInput{"attack", player[0], player, ""}

	Battle(userOneInput, userTwoInput)
}
