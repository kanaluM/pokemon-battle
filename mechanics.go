package main

import (
    "math"
    "math/rand"
)

// type multiplier function
// Takes the attacker's move and the defending pokemon 
// and calculates the type matchup damage multiplier
func TypeMatch(mv *Move, target *Pokemon) (float64, string) {
    moveType := mv.moveType   // attacking move's type

    defType1 := target.type1   // defending pokemon's first type
    defType1Num := Types[defType1]   // corresponding number of defType1
    tmult := Matchup[moveType][defType1Num]   // get multiplier from array

    defType2 := target.type2   // defending pokemon's second type if applicable
    if defType2 != "" {       // continue calculation if there is a second type
        defType2Num := Types[defType2]   // corresponding number of defType2
        tmult2 := Matchup[moveType][defType2Num]   // get multiplier from array
        tmult *= tmult2
    }

    // return type multiplier and a message about effectiveness
    message := ""
    if tmult < 1 { 
        message = "It's not very effective" 
    } else if tmult > 1 { 
        message = "It's super effective!" 
    }
    return tmult, message 
}


// turn order function
// takes in a list of pokemon in battle and returns
// a list of those pokemon ordered by speed (fast -> slow)
func TurnOrder(pkmnList [2]*Pokemon) ([2]*Pokemon) {
    if pkmnList[0].speed < pkmnList[1].speed {
        slow := pkmnList[0]
        pkmnList[0] = pkmnList[1]
        pkmnList[1] = slow
    } 

    return pkmnList
}


// takes a target pokemon and an attacking move and applies a volatile effect
// to the target if necessary. Returns the corresponding status message
// includes burn attack modifier and paralysis speed modifier
func ApplyVolatileStatus(target *Pokemon, move *Move) (string) {

    // does the target already have a non-volatile status effect
    // can only have one, so return if so
    if target.volatileStatus != "" { return "" }

    // does the move have a relevant secondary effect?
    effect := move.secondEffect     // the status that could occur
    _, isStatus := StatusList[effect]
    if !isStatus { return "" }      // if not then do nothing and return


    // does secondary effect actually happen?
    prob := move.secondEffectRate      // probability of status occuring
    threshold := rand.Float64()        // randomly generated threshold  0 <= x < 1
    if threshold > prob { return "" }  // if not then do nothing and return

    // apply status to the target and return message 
    target.volatileStatus = effect
    switch effect {
        case "PSN":
            return target.name + " was poisoned!"
        case "BRN":
            target.atk = target.atk / 2
            return target.name + " was burned!"
        case "FRZ":
            return target.name + " was frozen solid!"
        case "PRZ":
            target.speed = target.speed / 4
            return target.name + " was paralyzed!"
    }
    return ""
}


// takes a target pokemon and an attacking move and applies a non-volatile effect
// to the target if necessary. Returns the corresponding status message
func ApplyNonVolatileStatus(target *Pokemon, move *Move) (string) {
    
    // does the target already have a non-volatile status effect
    // can only have one, so return if so
    // if target.nonVolatileStatus != "" { return nil }

    // does the move have a relevant secondary effect?
    effect := move.secondEffect     // the status that could occur
    _, isStatus := StatusList[effect]
    if !isStatus { return "" }   // if not then do nothing and return

    // does secondary effect actually happen?
    prob := move.secondEffectRate      // probability of status occuring
    threshold := rand.Float64()        // randomly generated threshold  0 <= x < 1
    if threshold > prob { return "" }  // if not then do nothing and return

    // apply status to the target and return message 
    target.nonVolatileStatus = effect
    switch effect {
        case "flinch":
            return ""
    }
    return ""
}


// checks if status (volatile) prevents attacking this turn
// returns true if can attack, false if not. Also returns an appropriate status message
func CanAttackWithVolatileStatus(pokemon *Pokemon) (bool, string) {
    switch pokemon.volatileStatus {
        case "FRZ":
            // when frozen you cannot attack
            // each turn there is a 20% chance to thaw out
            if rand.Float64() <= 0.2 {
                pokemon.volatileStatus = ""
                return true, pokemon.name + " thawed out!"
            } else {
                return false, pokemon.name + " is frozen solid"
            }
        case "PRZ":
            // when paralyzed there is a 25% chance to be unable to attack
            if rand.Float64() <= 0.75 {
                return true, ""
            } else {
                return false, pokemon.name + " is paralyzed and can't move"
            }
        default:
            return true, ""
    }
}


// checks if status (volatile) prevents attacking this turn
// returns true if can attack, false if not. Also returns an appropriate status message
func CanAttackWithNonVolatileStatus(pokemon *Pokemon) (bool, string) {
    switch pokemon.nonVolatileStatus {
        case "flinch":
            // when flinched you cannot attack
            // lose flinch status at end of turn
            return false, pokemon.name + " flinched!"
        default:
            return true, ""
    }
}
   

// damage multiplier calculator
// returns multiplier and relevant messages
func DamageMultiplier(attacker *Pokemon, defender *Pokemon, move *Move) (float64, []string) {

    // check for critical hit
    isCrit := false
    if move.secondEffect == "crit" {   // high crit moves -> 1/8 chance
        if rand.Float64() <= 0.125 { 
            isCrit = true 
        }
    } else if rand.Float64() <= 0.0625 { 
        isCrit = true 
    }   // normal moves -> 1/16 chance

    // random damage roll
    var rngMult float64 = 0.85 + rand.Float64() * (1.0 - 0.85)
    
    // apply stab
    var stab float64 = 1
    if attacker.type1 == move.moveType || attacker.type2 == move.moveType {
        stab = 1.5
    }

    // move type vs defender type(s) multiplier
    typeMult, typeMessage := TypeMatch(move, defender)

    // return total multiplier
    var messages = []string{}
    if typeMessage != "" {
        messages = append(messages, typeMessage)
    }
    totalMult := rngMult * stab * typeMult

    if isCrit { 
        totalMult *= 2 
        messages = append(messages, "A critical hit!") 
    }

    return totalMult, messages
}

  
// accuracy check
// returns bool indicating if attack lands and a relevant message
func AccuracyCheck(pokemon *Pokemon, move *Move) (bool, string) {
    if rand.Intn(100) < move.accuracy {
        return true, ""
    } else {
        return false, pokemon.name + " missed!"
    }
}


// damage calculator
// returns damage dealt by an attack plus relevant messages
// does not do accuracy check or any other checks
func DamageCalc(attacker *Pokemon, defender *Pokemon, move *Move) (int, []string) {

    // either a special or physical interaction
    attackStat := attacker.atk
    defenseStat := defender.def
    if move.atkType == "spatk" {
        attackStat = attacker.spatk
        defenseStat = defender.spdef
    }

    // additional information for damage calculation
    level := attacker.level
    basePower := move.power
    multiplier, messages := DamageMultiplier(attacker, defender, move)
    damage := ((((2*float64(level)/5)+2)*float64(basePower)*(float64(attackStat)/float64(defenseStat))/50)+2)*multiplier
    return int(math.Round(damage * multiplier)), messages
}



