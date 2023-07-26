package main

import (
	"fmt"
)

func main() {
	attacker := newPokemon("Torterra", false)
	defender := newPokemon("Garchomp", false)
	fmt.Println(attacker.name + " vs. " + defender.name)

	userOneInput := UserInput{"attack", attacker, ""}
	userTwoInput := UserInput{"attack", attacker, ""}

	msgs := WholeTurn(userOneInput &UserInput,userTwoInput &UserInput)
	for x := range msgs[0] {
		fmt.Println(x)
	}
	for y := range msgs[1] {
		fmt.Println(y)
	}
}
