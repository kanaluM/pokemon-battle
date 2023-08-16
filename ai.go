package main

// Handles AI choosing a move
func ChooseMoveAI(aiInput *UserInput, userInput *UserInput) *UserInput {
	// AI chooses the move with the highest damage
	damage := 0
	var chosenMove string
	for _, mv := range aiInput.activePokemon.moves {
		move := MoveList[mv]
		dmg, _ := DamageCalc(aiInput.activePokemon, userInput.activePokemon, &move)
		if dmg > damage {
			damage = dmg
			chosenMove = mv
		}
	}
	aiInput.move = chosenMove
	return aiInput
}

// Handles AI sending out a new pokemon when needed
func ReplaceFaintedPokemonAI(aiInput *UserInput, userInput *UserInput) *UserInput {
	// AI sends out pokemon with strongest move
	damage := 0
	var chosenPokemon *Pokemon
	for _, pokemon := range aiInput.team {
		// don't consider fainted pokemon
		if pokemon.fainted {
			continue
		}
		// AI chooses the move with the highest damage
		for _, mv := range pokemon.moves {
			move := MoveList[mv]
			dmg, _ := DamageCalc(aiInput.activePokemon, userInput.activePokemon, &move)
			if dmg > damage {
				damage = dmg
				chosenPokemon = pokemon
			}
		}
	}
	
	// check if a pokemon to send out was found
	if chosenPokemon == nil {
		aiInput.gameOver = true
	} else {
		aiInput.activePokemon = chosenPokemon
	}
	return aiInput
}