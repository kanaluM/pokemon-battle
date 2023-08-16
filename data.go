package main

// list of non-volatile status conditions
var StatusList = map[string]bool{
    "PSN": true, 
    "FRZ": true, 
    "BRN": true, 
    "PRZ": true,
}


// List of pokemon types
var Types = map[string]int{
    "Normal": 0, 
    "Fire": 1, 
    "Water": 2, 
    "Electric": 3, 
    "Grass": 4, 
    "Ice": 5, 
    "Fighting": 6, 
    "Poison": 7, 
    "Ground": 8, 
    "Flying": 9, 
    "Psychic": 10, 
    "Bug": 11,
    "Rock": 12, 
    "Ghost": 13, 
    "Dragon": 14, 
    "Dark": 15, 
    "Steel": 16,
}


// type multipliers based on type
// KEY is attacking type
// VALUES are multiplier based on defending type
var Matchup = map[string][17]float64{
    "Normal": [17]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.5, 0, 1, 1, 0.5}, 
    "Fire": [17]float64{1, 0.5, 0.5, 1, 2, 2, 1, 1, 1, 1, 1, 2, 0.5, 1, 0.5, 1, 1},
    "Water": [17]float64{1, 2, 0.5, 1, 0.5, 1, 1, 1, 2, 1, 1, 1, 2, 1, 0.5, 1, 1},
    "Electric": [17]float64{1, 1, 2, 0.5, 0.5, 1, 1, 1, 0, 2, 1, 1, 1, 1, 0.5, 1, 1},
    "Grass": [17]float64{1, 0.5, 2, 1, 0.5, 1, 1, 0.5, 2, 0.5, 1, 0.5, 2, 1, 0.5, 1, 0.5},
    "Ice": [17]float64{1, 0.5, 0.5, 1, 2, 0.5, 1, 1, 2, 2, 1, 1, 1, 1, 2, 1, 0.5},
    "Fighting": [17]float64{2, 1, 1, 1, 1, 2, 1, 0.5, 1, 0.5, 0.5, 0.5, 2, 0, 1, 2, 2},
    "Poison": [17]float64{1, 1, 1, 1, 2, 1, 1, 0.5, 0.5, 1, 1, 1, 0.5, 0.5, 1, 1, 0},
    "Ground": [17]float64{1, 2, 1, 2, 0.5, 1, 1, 2, 1, 0, 1, 0.5, 2, 1, 1, 1, 2},
    "Flying": [17]float64{1, 1, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 2, 0.5, 1, 1, 1, 0.5},
    "Psychic": [17]float64{1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 0.5, 1, 1, 1, 1, 0, 0.5},
    "Bug": [17]float64{1, 0.5, 1, 1, 2, 1, 0.5, 0.5, 1, 0.5, 2, 1, 1, 0.5, 1, 2, 0.5},
    "Rock": [17]float64{1, 2, 1, 1, 1, 2, 0.5, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 0.5},
    "Ghost": [17]float64{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5},
    "Dragon": [17]float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0.5},
    "Dark": [17]float64{1, 1, 1, 1, 1, 1, 0.5, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5},
    "Steel": [17]float64{1, 0.5, 0.5, 0.5, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0.5},
}

// move, type, special/physical, power, accuracy, secondary effect
type Move struct {
    name string
    moveType string
    atkType string
    power int
    accuracy int
    secondEffectRate float64
    secondEffect string
}

// maps move name to information
var MoveList = map[string]Move{
    "Strength": Move{"Strength", "Normal", "atk", 80, 100, 0, "None"},
    "Hyper Voice": Move{"Hyper Voice", "Normal", "spatk", 90, 100, 0, "None"},
    "Flamethrower": Move{"Flamethrower", "Fire", "spatk", 95, 100, 0.1, "BRN"},
    "Flame Wheel": Move{"Flame Wheel", "Fire", "atk", 60, 100, 0.1, "BRN"},
    "Surf": Move{"Surf", "Water", "spatk", 95, 100, 0, "None"},
    "Waterfall": Move{"Waterfall", "Water", "atk", 80, 100, 0.2, "flinch"},
    "Thunderbolt": Move{"Thunderbolt", "Electric", "spatk", 95, 100, 0.1, "PRZ"},
    "Spark": Move{"Spark", "Electric", "atk", 65, 100, 0.3, "PRZ"},
    "Energy Ball": Move{"Energy Ball", "Grass", "spatk", 80, 100, 0.1, "spdef"},
    "Seed Bomb": Move{"Seed Bomb", "Grass", "atk", 80, 100, 0, "None"},
    "Leaf Blade": Move{"Leaf Blade", "Grass", "atk", 90, 100, 2, "crit"},
    "Ice Beam": Move{"Ice Beam", "Ice", "spatk", 95, 100, 0.1, "FRZ"},
    "Avalanche": Move{"Avalanche", "Ice", "atk", 80, 100, 0, "None"},
    "Aura Sphere": Move{"Aura Sphere", "Fighting", "spatk", 90, 100, 0, "None"},
    "Brick Break": Move{"Brick Break", "Fighting", "atk", 75, 100, 0, "None"},
    "Sludge Bomb": Move{"Sludge Bomb", "Poison", "spatk", 90, 100, 0.3, "PSN"},
    "Poison Jab": Move{"Poison Jab", "Poison", "atk", 90, 100, 0.3, "PSN"},
    "Earthquake": Move{"Earthquake", "Ground", "atk", 100, 100, 0, "None"},
    "Earth Power": Move{"Earth Power", "Ground", "spatk", 90, 100, 10, "spdef"},
    "Air Slash": Move{"Air Slash", "Flying", "spatk", 75, 95, 0.3, "flinch"},
    "Aerial Ace": Move{"Aerial Ace", "Flying", "atk", 60, 100, 0, "None"},
    "Psychic": Move{"Psychic", "Psychic", "spatk", 90, 100, 0.1, "spdef"},
    "Psycho Cut": Move{"Psycho Cut", "Psychic", "atk", 90, 100, 0, "None"},
    "Extrasensory": Move{"Extrasensory", "Psychic", "spatk", 80, 100, 0.1, "flinch"},
    "X-scissor": Move{"X-scissor", "Bug", "atk", 80, 100, 0, "None"},
    "Bug Buzz": Move{"Bug Buzz", "Bug", "spatk", 90, 100, 0.1, "spdef"},
    "Rock Slide": Move{"Rock Slide", "Rock", "atk", 75, 90, 0.3, "flinch"},
    "Stone Edge": Move{"Stone Edge", "Rock", "atk", 100, 80, 2, "crit"},
    "Power Gem": Move{"Power Gem", "Rock", "spatk", 80, 100, 0, "None"},
    "Shadow Ball": Move{"Shadow Ball", "Ghost", "spatk", 80, 100, 0.2, "spdef"},
    "Shadow Punch": Move{"Shadow Punch", "Ghost", "atk", 60, 100, 0, "None"},
    "Dragon Pulse": Move{"Dragon Pulse", "Dragon", "spatk", 90, 100, 0, "None"},
    "Dragon Claw": Move{"Dragon Claw", "Dragon", "atk", 80, 100, 0, "None"},
    "Dark Pulse": Move{"Dark Pulse", "Dark", "spatk", 80, 100, 0.2, "flinch"},
    "Night Slash": Move{"Night Slash", "Dark", "atk", 70, 100, 2, "crit"},
    "Crunch": Move{"Crunch", "Dark", "atk", 80, 100, 0.2, "def"},
    "Iron Head": Move{"Iron Head", "Steel", "atk", 80, 100, 0.3, "flinch"},
    "Flash Cannon": Move{"Flash Cannon", "Steel", "spatk", 80, 100, 0.1, "spdef"},
}

type PokemonData struct {
    // shared among individuals
    name string
    pokedexNumber int
    type1 string
    type2 string
    baseHp int
    baseAtk int
    baseDef int
    baseSpatk int
    baseSpdef int
    baseSpeed int
    moves [4]string
}

type Pokemon struct {
    // shared among individuals
    PokemonData

    // specific per individual
    level int
    hp int
    atk int
    def int
    spatk int
    spdef int
    speed int
    nonVolatileStatus string
    volatileStatus string
    fainted bool
}

var PokemonList = map[string]PokemonData{ 
    "Spiritomb": PokemonData{"Spiritomb", 108, "Ghost", "Dark", 50, 92, 108, 92, 108, 35, [4]string{"Dark Pulse", "Shadow Ball", "Bug Buzz", "Psychic"}}, 
    "Lucario": PokemonData{"Lucario", 116, "Fighting", "Steel", 70, 110, 70, 115, 70, 90, [4]string{"Strength", "Aura Sphere", "Shadow Ball", "Stone Edge"}}, 
    "Togekiss": PokemonData{"Togekiss", 175, "Normal", "Flying", 85, 50, 95, 120, 115, 80, [4]string{"Air Slash", "Aura Sphere", "Surf", "Thunderbolt"}}, 
    "Roserade": PokemonData{"Roserade", 27, "Grass", "Poison", 60, 70, 55, 125, 105, 90, [4]string{"Energy Ball", "Sludge Bomb", "Extrasensory", "Shadow Ball"}}, 
    "Milotic": PokemonData{"Milotic", 139, "Water", "", 95, 60, 79, 100, 125, 81, [4]string{"Surf", "Ice Beam", "Dragon Pulse", "Flamethrower"}}, 
    "Garchomp": PokemonData{"Garchomp", 111, "Dragon", "Ground", 108, 130, 95, 80, 85, 102, [4]string{"Dragon Claw", "Earthquake", "Flamethrower", "Rock Slide"}},
    "Torterra": PokemonData{"Torterra", 3, "Grass", "Ground", 95, 109, 105, 75, 85, 56, [4]string{"Energy Ball", "Earthquake", "Crunch", "Rock Slide"}}, 
    "Infernape": PokemonData{"Infernape", 6, "Fire", "Fighting", 76, 104, 71, 104, 71, 108, [4]string{"Flamethrower", "Brick Break", "Aura Sphere", "Poison Jab"}}, 
    "Empoleon": PokemonData{"Empoleon", 9, "Water", "Steel", 84, 86, 88, 111, 101, 60, [4]string{"Surf", "Flash Cannon", "Aerial Ace", "Strength"}}, 
    "Staraptor": PokemonData{"Staraptor", 12, "Normal", "Flying", 85, 120, 70, 50, 50, 100, [4]string{"Strength", "Aerial Ace", "Brick Break", "Leaf Blade"}}, 
    "Bibarel": PokemonData{"Bibarel", 14, "Normal", "Water", 79, 85, 60, 55, 60, 71, [4]string{"Strength", "Waterfall", "Aerial Ace", "Brick Break"}}, 
    "Luxray": PokemonData{"Luxray", 19, "Electric", "", 80, 120, 79, 95, 79, 70, [4]string{"Spark", "Crunch", "Brick Break", "Strength"}}, 
    "Alakazam": PokemonData{"Alakazam", 22, "Psychic", "", 55, 50, 45, 135, 85, 120, [4]string{"Psychic", "Shadow Ball", "Dark Pulse", "Aura Sphere"}}, 
    "Gyarados": PokemonData{"Gyarados", 24, "Water", "Flying", 95, 125, 79, 60, 100, 81, [4]string{"Waterfall", "Aerial Ace", "Earthquake", "Hyper Voice"}}, 
    "Steelix": PokemonData{"Steelix", 35, "Steel", "Ground", 75, 85, 200, 55, 65, 30, [4]string{"Iron Head", "Earthquake", "Strength", "Stone Edge"}}, 
    "Machamp": PokemonData{"Machamp", 42, "Fighting", "", 90, 130, 80, 65, 85, 55, [4]string{"Brick Break", "Rock Slide", "Strength", "Earthquake"}}, 
    "Gastrodon": PokemonData{"Gastrodon", 61, "Water", "Ground", 111, 83, 68, 92, 82, 39, [4]string{"Earth Power", "Surf", "Ice Beam", "Sludge Bomb"}}, 
    "Heracross": PokemonData{"Heracross", 62, "Bug", "Fighting", 80, 125, 75, 40, 95, 85, [4]string{"Brick Break", "X-scissor", "Earthquake", "Night Slash"}}, 
    "Gengar": PokemonData{"Gengar", 71, "Ghost", "Poison", 60, 65, 60, 130, 75, 110, [4]string{"Shadow Ball", "Sludge Bomb", "Dark Pulse", "Aura Sphere"}}, 
    "Honchkrow": PokemonData{"Honchkrow", 75, "Dark", "Flying", 100, 125, 52, 105, 52, 71, [4]string{"Night Slash", "Aerial Ace", "Shadow Ball", "Dark Pulse"}}, 
    "Rapidash": PokemonData{"Rapidash", 91, "Fire", "", 65, 100, 70, 80, 80, 105, [4]string{"Flame Wheel", "Poison Jab", "Strength", "Iron Head"}}, 
    "Pikachu": PokemonData{"Pikachu", 104, "Electric", "", 35, 55, 50, 50, 40, 90, [4]string{"Thunderbolt", "Flamethrower", "Surf", "Aura Sphere"}}, 
    "Snorlax": PokemonData{"Snorlax", 113, "Normal", "", 160, 110, 65, 65, 110, 30, [4]string{"Strength", "Rock Slide", "Crunch", "Earthquake"}}, 
    "Drapion": PokemonData{"Drapion", 128, "Poison", "Dark", 70, 90, 110, 60, 75, 95, [4]string{"Poison Jab", "Night Slash", "X-scissor", "Earthquake"}}, 
    "Toxicroak": PokemonData{"Toxicroak", 130, "Poison", "Fighting", 83, 106, 65, 86, 65, 85, [4]string{"Poison Jab", "Brick Break", "Extrasensory", "Night Slash"}}, 
    "Abomasnow": PokemonData{"Abomasnow", 143, "Grass", "Ice", 90, 92, 75, 92, 85, 60, [4]string{"Avalanche", "Seed Bomb", "Ice Beam", "Earthquake"}}, 
    "Weavile": PokemonData{"Weavile", 145, "Dark", "Ice", 70, 120, 65, 45, 85, 125, [4]string{"Night Slash", "Avalanche", "X-scissor", "Iron Head"}}, 
    "Gardevoir": PokemonData{"Gardevoir", 159, "Psychic", "", 68, 65, 65, 125, 115, 80, [4]string{"Psychic", "Shadow Ball", "Aura Sphere", "Energy Ball"}}, 
    "Gallade": PokemonData{"Gallade", 160, "Psychic", "Fighting", 68, 125, 65, 65, 115, 80, [4]string{"Leaf Blade", "Psycho Cut", "Brick Break", "Strength"}}, 
    "Altaria": PokemonData{"Altaria", 172, "Dragon", "Flying", 75, 70, 90, 70, 105, 80, [4]string{"Dragon Pulse", "Dragon Claw", "Aerial Ace", "Ice Beam"}}, 
    "Houndoom": PokemonData{"Houndoom", 177, "Dark", "Fire", 75, 90, 50, 110, 80, 95, [4]string{"Dark Pulse", "Flamethrower", "Energy Ball", "Earth Power"}}, 
    "Magnezone": PokemonData{"Magnezone", 180, "Electric", "Steel", 70, 70, 115, 130, 90, 60, [4]string{"Flash Cannon", "Thunderbolt", "Hyper Voice", "Bug Buzz"}}, 
    "Rhyperior": PokemonData{"Rhyperior", 188, "Ground", "Rock", 115, 140, 130, 55, 55, 40, [4]string{"Earthquake", "Stone Edge", "Poison Jab", "Brick Break"}}, 
    "Dusknoir": PokemonData{"Dusknoir", 191, "Ghost", "", 45, 100, 135, 65, 135, 45, [4]string{"Shadow Punch", "Night Slash", "Rock Slide", "Strength"}}, 
    "Porygon-z": PokemonData{"Porygon-z", 194, "Normal", "", 85, 80, 70, 135, 75, 90, [4]string{"Hyper Voice", "Thunderbolt", "Dark Pulse", "Ice Beam"}}, 
    "Scizor": PokemonData{"Scizor", 196, "Bug", "Steel", 70, 130, 100, 55, 80, 65, [4]string{"Iron Head", "X-scissor", "Strength", "Night Slash"}}, 
    "Electivire": PokemonData{"Electivire", 199, "Electric", "", 75, 123, 67, 95, 85, 95, [4]string{"Spark", "Earthquake", "Iron Head", "Rock Slide"}}, 
    "Magmortar": PokemonData{"Magmortar", 202, "Fire", "", 75, 95, 67, 125, 95, 83, [4]string{"Flamethrower", "Dark Pulse", "Thunderbolt", "Psychic"}}, 
    "Mamoswine": PokemonData{"Mamoswine", 205, "Ice", "Ground", 110, 130, 80, 70, 60, 80, [4]string{"Avalanche", "Earthquake", "Strength", "Stone Edge"}}, 
    "Froslass": PokemonData{"Froslass", 208, "Ice", "Ghost", 70, 80, 70, 80, 70, 110, [4]string{"Shadow Ball", "Ice Beam", "Thunderbolt", "Psychic"}},
    "MissingNo": PokemonData{"MissingNo", 0, "Dragon", "Ghost", 300, 300, 300, 300, 300, 300, [4]string{"Shadow Ball", "Dragon Pulse", "Thunderbolt", "Surf"}},
} 

var WinMessages = []string{"That was excellent. Truly, an outstanding battle. You gave the support your Pokémon needed to maximize their power. And you guided them with certainty to secure victory. You have both passion and calculating coolness. Together, you and your Pokémon can overcome any challenge that may come your way. Those are the impressions I got from our battle. I'm glad I got to take part in the crowning of Sinnoh's new Champion! Come with me. We'll take the lift."}   // win messages
var LoseMessages = []string{"Smell ya later!", "Better luck next time", "Keep training", "Time to soft-reset", "You whited out...", "Come back when you're stronger", "Do or do not...there is no try", "Looks like you're blasting off again"}   // lose messages

type UserInput struct {
    username string
    action string
    activePokemon *Pokemon
    team []*Pokemon
    move string
    isAI bool
    gameOver bool
}