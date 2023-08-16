package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

// the turn function for player
// returns a list of messages
// checks for status hindrance, accuracy
func AttackTurn(attacker *Pokemon, defender *Pokemon, move *Move) ([]string) {

	// check for volatile status
	canAttack, msg := CanAttackWithVolatileStatus(attacker) 
	if !canAttack {
		return []string{msg}
	}

	// check for non-volatile status
	canAttack, msg = CanAttackWithNonVolatileStatus(attacker) 
	if !canAttack {
		return []string{msg}
	}

	// initialize slice of messages to return
	messages := []string{attacker.name + " used " + move.name}

	// accuracy check
    attackLands, message := AccuracyCheck(attacker, move)
    if !attackLands {
        return append(messages, message)
    }

	// move lands, calculate damage dealt
	damage, msgs := DamageCalc(attacker, defender, move)
	messages = append(messages, msgs...)

	// apply damage to target
	if defender.hp <= damage {   // target faints
		defender.hp = 0
		defender.fainted = true
		messages = append(messages, defender.name + " fainted!")
	} else {   // target survives
		defender.hp -= damage
		messages = append(messages, defender.name + " has " + strconv.Itoa(defender.hp) + " hp left")
	}
            
	// apply statuses
	msg = ApplyVolatileStatus(defender, move)
	if msg != ""{ 
		messages = append(messages, msg) 
	}
	msg = ApplyNonVolatileStatus(defender, move)
	if msg != "" { 
		messages = append(messages, msg) 
	}

	return messages
}


// after each turn, check all pokemon on field for status effects and apply damage
// func PostTurnStatusEffects(pokemonInPlay [2]*Pokemon) ([]string) {
// 	pokemonInPlay = TurnOrder(pokemonInPlay)
// 	var messages []string

// 	// check each pokemon for volatile status effects
// 	for _, pokemon := range pokemonInPlay {

// 		// check for residual status damage
// 		damage := 0
// 		switch pokemon.volatileStatus {
// 			case "BRN": 
// 				messages = append(messages, pokemon.name + " was hurt by burn")
// 				damage = 8
// 			case "PSN": 
// 				messages = append(messages, pokemon.name + " was hurt by poison")
// 				damage = 16
// 			default:
// 				continue
// 		}

// 		// apply damage to target if needed
// 		if pokemon.hp <= damage {   // target faints
// 			pokemon.hp = 0
// 			pokemon.fainted = true
// 			messages = append(messages, pokemon.name + " fainted!")
// 		} else {   // target survives
// 			pokemon.hp -= damage
// 		}
// 	}

// 	return messages
// }


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
	if wonBattle {
		return winMessages[rand.Intn(len(winMessages))]
	} else {
		return loseMessages[rand.Intn(len(loseMessages))]
	}
}


// wrapper function for everything that happens in one turn
// returns a bunch of messages for each player
func WholeTurn(userOneInput *UserInput, userTwoInput *UserInput) ([2][]string, bool) {

	var msgs [2][]string
	var msg []string

	// turn order
	userOneInput, userTwoInput = TurnOrder(userOneInput, userTwoInput)
	// fmt.Println(userOneInput.activePokemon.name, userOneInput.activePokemon.speed)
	// fmt.Println(userTwoInput.activePokemon.name, userTwoInput.activePokemon.speed)

	// player 1 attacks
	// TODO cannot flinch when moving first (probably should be refactored in future)
	if userOneInput.activePokemon.nonVolatileStatus == "flinch" { 
		userOneInput.activePokemon.nonVolatileStatus = "" 
	}

	// faster user attacks or switches (AI only attacks)
	if userOneInput.isAI {
		userOneInput = ChooseMoveAI(userOneInput, userTwoInput)
		move := MoveList[userOneInput.move]
		msg = AttackTurn(userOneInput.activePokemon, userTwoInput.activePokemon, &move)
	} else {
		userOneInput = ChooseAction(userOneInput)
		if userOneInput.action == "attack" {
			move := MoveList[userOneInput.move]
			msg = AttackTurn(userOneInput.activePokemon, userTwoInput.activePokemon, &move)
		}
	}
	msgs[0] = msg

	// check if slower user has not fainted
	// if it has, then send out a new Pokemon
	if userTwoInput.activePokemon.fainted {
		if userTwoInput.isAI {
			userTwoInput = ReplaceFaintedPokemonAI(userTwoInput, userOneInput)
		} else {
			userTwoInput = ReplaceFaintedPokemon(userTwoInput)
		}
		if userTwoInput.gameOver {
			msgs[0] = append(msgs[0], "Someone is out of usable Pokemon...", "Someone whited out!") 
			return msgs, true
		}
	}

	// slower user attacks or switches (AI only attacks)
	if userTwoInput.isAI {
		userTwoInput = ChooseMoveAI(userTwoInput, userOneInput)
		move := MoveList[userTwoInput.move]
		msg = AttackTurn(userTwoInput.activePokemon, userOneInput.activePokemon, &move)
	} else {
		userTwoInput = ChooseAction(userTwoInput)
		if userTwoInput.action == "attack" {
			move := MoveList[userTwoInput.move]
			msg = AttackTurn(userTwoInput.activePokemon, userOneInput.activePokemon, &move)
		}
	}
	msgs[1] = msg

	// check if faster user has not fainted
	// if it has, then send out a new Pokemon
	if userOneInput.activePokemon.fainted {
		if userOneInput.isAI {
			userTwoInput = ReplaceFaintedPokemonAI(userOneInput, userTwoInput)
		} else {
			userOneInput = ReplaceFaintedPokemon(userOneInput)
		}
		if userOneInput.gameOver {
			msgs[1] = append(msgs[1], "Someone is out of usable Pokemon...", "Someone whited out!") 
			return msgs, true
		}
	}
	return msgs, false
}


// returns true if all pokemon on team are fainted
// may not be needed
func IsLoser(team []*Pokemon) bool {
	for _, p := range team {
		if !p.fainted {
			return false
		}
	}
	return true
}

// wrapper function for a whole 6v6 singles battle
func Battle(userOneInput *UserInput, userTwoInput *UserInput) {
	fmt.Println("[[ BATTLE ]] Starting a battle\n")

	var msgs [2][]string
	var gameOver bool
	for {
		msgs, gameOver = WholeTurn(userOneInput, userTwoInput)
		for _, x := range msgs[0] {
			fmt.Println(x)
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
		for _, y := range msgs[1] {
			fmt.Println(y)
			time.Sleep(1 * time.Second)
		}
		fmt.Println()
		if gameOver {
			break
		}
	}
}


