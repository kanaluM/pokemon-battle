package main

import (
	"fmt"
	"strconv"
)

func ChooseAction(input *UserInput) *UserInput {
	fmt.Println("Type a number and hit ENTR to choose an action:")
	fmt.Println("(1) Fight")
	fmt.Println("(2) Pokemon")
	fmt.Println("(3) Bag")
	fmt.Println("(4) Run")
	fmt.Println()
	var mv string
	canContinue := true
	for canContinue {
		fmt.Scanln(&mv)
		switch mv {
			case "1":
				canContinue = false
				fmt.Println()
			case "2":
				canContinue = false
				fmt.Println()
			case "3":
				fmt.Println("Your bag is empty")
			case "4":
				fmt.Println("Only cowards run away")
			default:
				fmt.Println("[[ INVALID INPUT ]] Try again")
		}
	}
	action, _ := strconv.Atoi(mv)
	res := input
	switch action {
		case 1: 
			input.action = "attack"
			res = ChooseMove(input)
			fmt.Println()
		case 2:
			input.action = "switch"
			res = SwitchPokemon(input)
			fmt.Println()
		default:
			fmt.Println("huh")
	}
	return res
}

// gets user stdin to decide which move to use
func ChooseMove(input *UserInput) *UserInput {
	pokemon := input.activePokemon
	fmt.Println("Type a number and hit ENTR to choose a move:")
	for i := 0; i < 4; i++ {
		fmt.Printf("(%v) %v\n", i+1, pokemon.moves[i])
	}
	fmt.Println()
	var mv string
	for {
		fmt.Scanln(&mv)
		if (mv == "1") || (mv == "2") || (mv == "3") || (mv == "4") { break } 
		fmt.Println("[[ INVALID INPUT ]] Try again")
	}
	moveIdx, _ := strconv.Atoi(mv)
	// fmt.Println(pokemon.name, "used", pokemon.moves[moveIdx-1])
	input.move = pokemon.moves[moveIdx-1]
	return input
}

// gets user stdin to decide which Pokemon to send out 
// (after previous one has fainted)
func ReplaceFaintedPokemon(input *UserInput) *UserInput {
	team := input.team
	fmt.Println("Type a number and hit ENTR to choose a pokemon:")
	var nonFaintedPokemon []*Pokemon
	for i := 0; i < len(team); i++ {
		if !team[i].fainted {
			nonFaintedPokemon = append(nonFaintedPokemon, team[i])
			fmt.Printf("(%v) %v\n", len(nonFaintedPokemon), team[i].name)
		}
	}
	fmt.Println()

	// Are there more pokemon to fight?
	if len(nonFaintedPokemon) == 0 {
		input.gameOver = true
		return input
	}

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
	input.activePokemon = nonFaintedPokemon[res]
	return input
}

// gets user stdin to decide which Pokemon to send out 
// (switching during turn)
func SwitchPokemon(input *UserInput) *UserInput {
	activePokemon := input.activePokemon
	team := input.team
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
		return input
	}
	fmt.Println("\nCome back", activePokemon.name)
	fmt.Println("Go", nonFaintedPokemon[res].name)
	input.activePokemon = nonFaintedPokemon[res]
	return input
}