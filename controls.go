package main

import (
	"fmt"
	"strconv"
)

// gets user stdin to decide which move to use
func ChooseMove(pokemon *Pokemon) Move {
	fmt.Println("Type a number and hit ENTR to choose a move:")
	for i := 0; i < 4; i++ {
		fmt.Printf("(%v) %v\n", i+1, pokemon.moves[i])
	}
	fmt.Println()
	var mv string
	for {
		fmt.Scanln(&mv)
		if (mv == "1") || (mv == "2") || (mv == "3") || (mv == "4") { break } 
		fmt.Println("Try again bro")
	}
	moveIdx, _ := strconv.Atoi(mv)
	fmt.Println(pokemon.name, "used", pokemon.moves[moveIdx-1])
	return MoveList[mv]
}

// gets user stdin to decide which Pokemon to send out 
// (after previous one has fainted)
func ReplaceFaintedPokemon(team []*Pokemon) *Pokemon {
	fmt.Println("Type a number and hit ENTR to choose a pokemon:")
	var nonFaintedPokemon []*Pokemon
	for i := 0; i < len(team); i++ {
		if !team[i].fainted {
			nonFaintedPokemon = append(nonFaintedPokemon, team[i])
			fmt.Printf("(%v) %v\n", len(nonFaintedPokemon), team[i].name)
		}
	}
	fmt.Println()
	var mv string
	var res int
	for {
		fmt.Scanln(&mv)
		pokemonIdx, err := strconv.Atoi(mv)
		if err == nil && (0 < pokemonIdx) && (pokemonIdx < len(nonFaintedPokemon)+1) { 
			res = pokemonIdx-1
			break 
		}
		fmt.Println("[[ INVALID INPUT ]] Try again")
	}
	fmt.Println("\nSent out", nonFaintedPokemon[res].name)
	return nonFaintedPokemon[res]
}


// gets user stdin to decide which Pokemon to send out 
// (switching during turn)
func SwitchPokemon(team []*Pokemon, activePokemon *Pokemon) *Pokemon {
	fmt.Println("Type a number and hit ENTR to choose a pokemon:")
	var nonFaintedPokemon []*Pokemon
	fmt.Println("(0) [[ GO BACK ]]")
	for i := 0; i < len(team); i++ {
		if !team[i].fainted && team[i] != activePokemon {
			nonFaintedPokemon = append(nonFaintedPokemon, team[i])
			fmt.Printf("(%v) %v\n", len(nonFaintedPokemon), team[i].name)
		}
	}
	fmt.Println()
	var mv string
	var res int
	for {
		fmt.Scanln(&mv)
		pokemonIdx, err := strconv.Atoi(mv)
		if err == nil && (0 <= pokemonIdx) && (pokemonIdx < len(nonFaintedPokemon)+1) { 
			res = pokemonIdx-1
			break 
		}
		fmt.Println("[[ INVALID INPUT ]] Try again")
	}
	if res == -1 {
		return activePokemon
	}
	fmt.Println("\nCome back", activePokemon.name)
	fmt.Println("Go", nonFaintedPokemon[res].name)
	return nonFaintedPokemon[res]
}