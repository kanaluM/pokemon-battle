package main

import (
	// "fmt"
)

func main() {
	empoleon := NewPokemon("Empoleon", false)
	torterra := NewPokemon("Torterra", false)
	infernape := NewPokemon("Infernape", false)
	staraptor := NewPokemon("Staraptor", false)
	luxray := NewPokemon("Luxray", false)
	machamp := NewPokemon("Machamp", false)
	
	spiritomb := NewPokemon("Spiritomb", true)
	lucario := NewPokemon("Lucario", true)
	togekiss := NewPokemon("Togekiss", true)
	roserade := NewPokemon("Roserade", true)
	milotic := NewPokemon("Milotic", true)
	garchomp := NewPokemon("Garchomp", true)

	myTeam := []*Pokemon{empoleon, torterra, infernape, staraptor, luxray, machamp}
	cynthiasTeam := []*Pokemon{spiritomb, lucario, togekiss, roserade, milotic, garchomp}

	myInput := &UserInput{"", empoleon, myTeam, "", false, false}
	cynthiasInput := &UserInput{"", spiritomb, cynthiasTeam, "", true, false}

	Battle(myInput, cynthiasInput)
}
