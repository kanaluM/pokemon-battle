package main

import "math/rand"

// does math calculation for a pokemon's real HP stat
func calculateHp(pokemon *Pokemon, iv int, ev int) (int) {
    return (((2 * pokemon.baseHp) + iv + (ev / 4)) * pokemon.level) / 100 + pokemon.level + 10
}


// does math calculation for a pokemon's real HP stat
func calculateOtherStat(pokemon *Pokemon, iv int, ev int, baseStat int) (int) {
    return (((2 * baseStat) + iv + (ev / 4)) * pokemon.level) / 100 + 5
}


// stat calculator
// takes a pokemon and a bool indicating whether to make it strong or not
// and fills in that pokemon's actual stats (based on base stats, IVs, etc) 
func initializeStats(pokemon *Pokemon, makeStrong bool) {
    if makeStrong {   // cynthia AI will get stronger pokemon
        pokemon.level = 60
        IV := 31
        EV := 252
        pokemon.hp = calculateHp(pokemon, IV, EV)
        pokemon.atk = calculateOtherStat(pokemon, IV, EV, pokemon.baseAtk)
        pokemon.def = calculateOtherStat(pokemon, IV, EV, pokemon.baseDef)
        pokemon.spatk = calculateOtherStat(pokemon, IV, EV, pokemon.baseSpatk)
        pokemon.spdef = calculateOtherStat(pokemon, IV, EV, pokemon.baseSpdef)
        pokemon.speed = calculateOtherStat(pokemon, IV, EV, pokemon.baseSpeed)
    } else {   // player gets slightly weaker pokemon
        pokemon.level = rand.Intn(10)+50
        pokemon.hp = calculateHp(pokemon, rand.Intn(20)+10, rand.Intn(100)+120)
        pokemon.atk = calculateOtherStat(pokemon, rand.Intn(20)+10, rand.Intn(100)+120, pokemon.baseAtk)
        pokemon.def = calculateOtherStat(pokemon, rand.Intn(20)+10, rand.Intn(100)+120, pokemon.baseDef)
        pokemon.spatk = calculateOtherStat(pokemon, rand.Intn(20)+10, rand.Intn(100)+120, pokemon.baseSpatk)
        pokemon.spdef = calculateOtherStat(pokemon, rand.Intn(20)+10, rand.Intn(100)+120, pokemon.baseSpdef)
        pokemon.speed = calculateOtherStat(pokemon, rand.Intn(20)+10, rand.Intn(100)+120, pokemon.baseSpeed)
    }
}


// new pokemon constructor
// didn't want to use packages so it is very raw
func NewPokemon(name string, makeStrong bool) *Pokemon {
	template := PokemonList[name]
	pokemon := Pokemon{}

	// shared among individuals
    pokemon.name = template.name
    pokemon.pokedexNumber = template.pokedexNumber
    pokemon.type1 = template.type1
    pokemon.type2 = template.type2
    pokemon.baseHp = template.baseHp
    pokemon.baseAtk = template.baseAtk
    pokemon.baseDef = template.baseDef
    pokemon.baseSpatk = template.baseSpatk
    pokemon.baseSpdef = template.baseSpdef
    pokemon.baseSpeed = template.baseSpeed
    pokemon.moves = template.moves

    // level and all stats handled by helper
    initializeStats(&pokemon, makeStrong)

	// remaining battle specific fields
    pokemon.nonVolatileStatus = ""
    pokemon.volatileStatus = ""
    pokemon.fainted = false

	return &pokemon
}