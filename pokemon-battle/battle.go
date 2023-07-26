package main

// the turn function for player
// returns a list of messages
// checks for status hindrance, accuracy
func UserTurnAttack(attacker &Pokemon, defender &Pokemon, move &Move) ([]string) {

	// check for volatile status
	canAttack, msg := CanAttackWithVolatileStatus(attacker &Pokemon) 
	if !canAttack {
		return []string{msg}
	}

	// check for non-volatile status
	canAttack, msg = CanAttackWithNonVolatileStatus(attacker &Pokemon) 
	if !canAttack {
		return []string{msg}
	}

	// initialize slice of messages to return
	messages = []string{attacker.name + " used " + move.moveName}

	// accuracy check
    attackLands, message = AccuracyCheck(attacker, move)
    if !attackLands {
        return append(messages, message)
    }

	// move lands, calculate damage dealt
	damage, msgs = DamageCalc(attacker, defender, move)
	messages = append(messages, msgs...)

	// apply damage to target
	if defender.hp <= damage {   // target faints
		defender.hp = 0
		defender.fainted = true
		messages = append(messages, defender.name + " fainted!")
	} else {   // target survives
		defender.hp -= damage
	}
            
	// apply statuses
	msg = ApplyVolatileStatus(defender, move)
	if msg { messages = append(messages, msg) }
	msg = ApplyNonVolatileStatus(defender, move)
	if msg { messages = append(messages, msg) }

	return messages
}


// the turn function for AI opponent
// returns a list of messages
// checks for status hindrance, accuracy
func AITurnAttack(attacker &Pokemon, defender &Pokemon) ([]string) {

	// AI chooses the move with the highest damage
	damage := 0
	chosenMove := &Move{}
	for mv := attacker.moves {
		move := &moveList[mv]
		dmg, _ := DamageCalc(attacker, defender, move)
		if dmg > damage {
			damage = dmg
			chosenMove = move
		}
	}

	// same workflow as a user attacking
	return UserTurnAttack(attacker, defender, chosenMove)
}


// after each turn, check all pokemon on field for status effects and apply damage
func PostTurnStatusEffects(pokemonInPlay [2]&Pokemon) ([]string) {
	pokemonInPlay = TurnOrder(pokemonInPlay)
	messages := []string

	// check each pokemon for volatile status effects
	for pokemon := range pokemonInPlay {

		// check for residual status damage
		damage := 0
		switch pokemon.volatileStatus {
			case "BRN": 
				messages = append(messages, pokemon.name + " was hurt by burn")
				damage = 8
			case "PSN": 
				messages = append(messages, pokemon.name + " was hurt by poison")
				damage = 16
			default:
				continue
		}

		// apply damage to target if needed
		if pokemon.hp <= damage {   // target faints
			pokemon.hp = 0
			pokemon.fainted = true
			messages = append(messages, pokemon.name + " fainted!")
		} else {   // target survives
			pokemon.hp -= damage
		}
	}

	return messages
}


// simple message for win or loss
func BattleOutcomeMessage(wonBattle bool) (string) {
	if wonBattle {
		return "You defeated Cynthia!"
	} else {
		return "You lost to Cynthia..."
	}
}


// returns a more complex message based on the outcome of the battle
func PostBattleMessage(winMessages []string, loseMessages []string, wonBattle bool) (string) {
	reasons[rand.Intn(len(reasons))]
	if wonBattle {
		return winMessages[rand.Intn(len(winMessages))]
	} else {
		return loseMessages[rand.Intn(len(loseMessages))]
	}
}


// wrapper function for everything that happens in one turn
// returns a bunch of messages for each player
func WholeTurn(userOneInput &UserInput,userTwoInput &UserInput) ([2][]string) {

	var msgs [2][]string

	// turn order
	pkmn1 = userOneInput.pokemon
	pkmn2 = userTwoInput.pokemon
	var pkmnList [2]&Pokemon{pkmn1, pkmn2}
	pkmnList = TurnOrder(pkmnList)

	// player 1 attacks
	// TODO cannot flinch when moving first (probably should be refactored in future)
	if pkmn1.nonVolatileStatus == "flinch" { pkmn1.nonVolatileStatus = "" }
	msg := AITurnAttack(pkmn1, pkmn2)
	msgs[0] = msg

	// check if player 2 is still alive
	if pkmn2.fainted {
		// switch
	}

	// player 2 attacks
	msg = AITurnAttack(pkmn2, pkmn1)
	msgs[1] = msg

	// check if player 1 is still alive
	if pkmn1.fainted {
		// switch
	}

	return msgs
}
