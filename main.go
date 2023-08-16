package main

import (
	// "fmt"
)

func main() {
	missingno := NewPokemon("MissingNo", true)
	// empoleon := NewPokemon("Empoleon", true)
	torterra := NewPokemon("Torterra", true)
	infernape := NewPokemon("Infernape", true)
	staraptor := NewPokemon("Staraptor", true)
	luxray := NewPokemon("Luxray", true)
	machamp := NewPokemon("Machamp", true)
	
	spiritomb := NewPokemon("Spiritomb", true)
	lucario := NewPokemon("Lucario", true)
	togekiss := NewPokemon("Togekiss", true)
	roserade := NewPokemon("Roserade", true)
	milotic := NewPokemon("Milotic", true)
	garchomp := NewPokemon("Garchomp", true)

	myTeam := []*Pokemon{missingno, torterra, infernape, staraptor, luxray, machamp}
	cynthiasTeam := []*Pokemon{spiritomb, lucario, togekiss, roserade, milotic, garchomp}

	// type UserInput struct {
	// 	action string
	// 	activePokemon *Pokemon
	// 	team []*Pokemon
	// 	move string
	// 	isAI bool
	// 	gameOver bool
	// }
	myInput := &UserInput{"Ash", "", missingno, myTeam, "", false, false}
	cynthiasInput := &UserInput{"Cynthia", "", spiritomb, cynthiasTeam, "", true, false}

	Battle(myInput, cynthiasInput)
}
