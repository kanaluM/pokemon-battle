package main

import (
	"fmt"
)

func main() {
	spiritomb := NewPokemon("Spiritomb", true)
	lucario := NewPokemon("Lucario", true)
	togekiss := NewPokemon("Togekiss", true)
	roserade := NewPokemon("Roserade", true)
	milotic := NewPokemon("Milotic", true)
	garchomp := NewPokemon("Garchomp", true)

	cynthiasTeam := []*Pokemon{spiritomb, lucario, togekiss, roserade, milotic, garchomp}

	// type UserInput struct {
	// 	action string
	// 	activePokemon *Pokemon
	// 	team []*Pokemon
	// 	move string
	// 	isAI bool
	// 	gameOver bool
	// }
	myInput := &UserInput{"Ash", "", nil, nil, "", false, false}
	cynthiasInput := &UserInput{"Cynthia", "", spiritomb, cynthiasTeam, "", true, false}

	fmt.Println()
	ChooseName(myInput)
	ChooseTeam(myInput)
	myInput.activePokemon = myInput.team[0]

	Battle(myInput, cynthiasInput)
}
