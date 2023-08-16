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
		messages = append(messages, defender.name + " lost " + strconv.Itoa(defender.hp) + " health")
		defender.hp = 0
		defender.fainted = true
		messages = append(messages, defender.name + " fainted!")
		return messages
	} else {   // target survives
		defender.hp -= damage
		messages = append(messages, defender.name + " lost " + strconv.Itoa(damage) + " health (" + strconv.Itoa(defender.hp) + " hp left)")
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

// after each turn, check a pokemon on field for status effects and apply damage
// func PostTurnStatusEffects(input *UserInput) ([]string) {
// 	var messages []string
// 	pokemon := input.activePokemon

// 	// check for residual status damage
// 	damage := 0
// 	switch pokemon.volatileStatus {
// 		case "BRN": 
// 			messages = append(messages, pokemon.name + " was hurt by burn")
// 			damage = 8
// 		case "PSN": 
// 			messages = append(messages, pokemon.name + " was hurt by poison")
// 			damage = 16
// 		default:
// 			continue
// 	}

// 	// apply damage to target if needed
// 	if pokemon.hp <= damage {   // target faints
// 		pokemon.hp = 0
// 		pokemon.fainted = true
// 		messages = append(messages, pokemon.name + " fainted!")
// 	} else {   // target survives
// 		pokemon.hp -= damage
// 	}

// 	return messages
// }

// returns a more complex message based on the outcome of the battle
func PostBattleMessage(winner *UserInput, loser *UserInput, wonBattle bool) {
	fmt.Println(winner.username, "defeated", loser.username)
	if wonBattle {   // you won the battle
		fmt.Println(WinMessages[rand.Intn(len(WinMessages))], "\n")
	} else {        // you lost the battle
		fmt.Println(LoseMessages[rand.Intn(len(LoseMessages))], "\n")
	}
}

// Printing helper function
func PrintMessages(msgs []string) {
	if len(msgs) == 0 {
		return
	}
	for _, x := range msgs {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

// runs a user's input after it has been collected for the turn
func HalfTurn(attackerInput *UserInput, defenderInput *UserInput) (*UserInput, *UserInput, bool) {
	var msg []string
	if attackerInput.action == "attack" {
		move := MoveList[attackerInput.move]
		msg = AttackTurn(attackerInput.activePokemon, defenderInput.activePokemon, &move)
	}
	PrintMessages(msg)

	// if defender faints, then send out a new Pokemon
	defenderLives := true
	if defenderInput.activePokemon.fainted {
		defenderLives= false
		if defenderInput.isAI {
			defenderInput = ReplaceFaintedPokemonAI(defenderInput, attackerInput)
		} else {
			defenderInput = ReplaceFaintedPokemon(defenderInput)
		}
	}
	return attackerInput, defenderInput, defenderLives
}

// wrapper function for everything that happens in one turn
// returns a bunch of messages for each player
func WholeTurn(userOneInput *UserInput, userTwoInput *UserInput) bool {

	var canAttack bool

	fmt.Println("[[ NEW TURN ]] What will you do?\n")

	// turn order
	userOneInput, userTwoInput = TurnOrder(userOneInput, userTwoInput)

	// TODO cannot flinch when moving first (probably should be refactored in future)
	if userOneInput.activePokemon.nonVolatileStatus == "flinch" { 
		userOneInput.activePokemon.nonVolatileStatus = "" 
	}

	// We want to collect the input from both users before anything happens in the turn
	// faster user attacks or switches (AI only attacks)
	if userOneInput.isAI {
		userOneInput = ChooseActionAI(userOneInput, userTwoInput)
	} else {
		userOneInput = ChooseAction(userOneInput)
	}

	// slower user attacks or switches (AI only attacks)
	if userTwoInput.isAI {
		userTwoInput = ChooseActionAI(userTwoInput, userOneInput)
	} else {
		userTwoInput = ChooseAction(userTwoInput)
	}

	// After collecting the input from both players, the turn can proceed
	// Start with faster pokemon attacking (or switching)
	userOneInput, userTwoInput, canAttack = HalfTurn(userOneInput, userTwoInput)

	// slower pokemon does not get a turn if fainted
	if userTwoInput.gameOver { 
		fmt.Println(userTwoInput.username, "is out of usable Pokemon...")
		fmt.Println(userTwoInput.username, "whited out!\n") 
		return true
	} else if !canAttack {
		return false
	}
	
	// slower pokemon can attack (or switch)
	userTwoInput, userOneInput, _ = HalfTurn(userTwoInput, userOneInput)
	
	if userOneInput.gameOver {
		fmt.Println(userOneInput.username, "is out of usable Pokemon...")
		fmt.Println(userOneInput.username, "whited out!\n") 
		return true
	} 
	return false
}

// wrapper function for a whole 6v6 singles battle
func Battle(userOneInput *UserInput, userTwoInput *UserInput) {
	fmt.Println("\n[[ BATTLE ]] Starting a battle\n")
	fmt.Println(userOneInput.username, "sent out", userOneInput.activePokemon.name)
	fmt.Println(userTwoInput.username, "sent out", userTwoInput.activePokemon.name, "\n")

	// var msgs [2][]string
	var gameOver bool
	for {
		gameOver = WholeTurn(userOneInput, userTwoInput)
		if gameOver {
			if userOneInput.gameOver {
				PostBattleMessage(userTwoInput, userOneInput, false)
			} else {
				PostBattleMessage(userOneInput, userTwoInput, true)
			}
			break
		}
	}
}


