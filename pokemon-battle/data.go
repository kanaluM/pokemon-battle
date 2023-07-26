package main

// list of non-volatile status conditions
const StatusList map[string]bool{
    "PSN": true, 
    "FRZ": true, 
    "BRN": true, 
    "PRZ": true
}


// List of pokemon types
const Types map[string]int{
    "Normal": 1, 
    "Fire": 2, 
    "Water": 3, 
    "Electric": 4, 
    "Grass": 5, 
    "Ice": 6, 
    "Fighting": 7, 
    "Poison": 8, 
    "Ground": 9, 
    "Flying": 10, 
    "Psychic": 11, 
    "Bug": 12,
    "Rock": 13, 
    "Ghost": 14, 
    "Dragon": 15, 
    "Dark": 16, 
    "Steel": 17
}


// type multipliers based on type
// KEY is attacking type
// VALUES are multiplier based on defending type
const Matchup map[string][17]float{
    "Normal": [17]float{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0.5, 0, 1, 1, 0.5}, 
    "Fire": [17]float{1, 0.5, 0.5, 1, 2, 2, 1, 1, 1, 1, 1, 2, 0.5, 1, 0.5, 1, 1},
    "Water": [17]float{1, 2, 0.5, 1, 0.5, 1, 1, 1, 2, 1, 1, 1, 2, 1, 0.5, 1, 1},
    "Electric": [17]float{1, 1, 2, 0.5, 0.5, 1, 1, 1, 0, 2, 1, 1, 1, 1, 0.5, 1, 1},
    "Grass": [17]float{1, 0.5, 2, 1, 0.5, 1, 1, 0.5, 2, 0.5, 1, 0.5, 2, 1, 0.5, 1, 0.5},
    "Ice": [17]float{1, 0.5, 0.5, 1, 2, 0.5, 1, 1, 2, 2, 1, 1, 1, 1, 2, 1, 0.5},
    "Fighting": [17]float{2, 1, 1, 1, 1, 2, 1, 0.5, 1, 0.5, 0.5, 0.5, 2, 0, 1, 2, 2},
    "Poison": [17]float{1, 1, 1, 1, 2, 1, 1, 0.5, 0.5, 1, 1, 1, 0.5, 0.5, 1, 1, 0},
    "Ground": [17]float{1, 2, 1, 2, 0.5, 1, 1, 2, 1, 0, 1, 0.5, 2, 1, 1, 1, 2},
    "Flying": [17]float{1, 1, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 2, 0.5, 1, 1, 1, 0.5},
    "Psychic": [17]float{1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 0.5, 1, 1, 1, 1, 0, 0.5},
    "Bug": [17]float{1, 0.5, 1, 1, 2, 1, 0.5, 0.5, 1, 0.5, 2, 1, 1, 0.5, 1, 2, 0.5},
    "Rock": [17]float{1, 2, 1, 1, 1, 2, 0.5, 1, 0.5, 2, 1, 2, 1, 1, 1, 1, 0.5},
    "Ghost": [17]float{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5},
    "Dragon": [17]float{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 0.5},
    "Dark": [17]float{1, 1, 1, 1, 1, 1, 0.5, 1, 1, 1, 2, 1, 1, 2, 1, 0.5, 0.5},
    "Steel": [17]float{1, 0.5, 0.5, 0.5, 1, 2, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 0.5}
}

// move, type, special/physical, power, accuracy, secondary effect
type Move struct {
    moveName string
    moveType string
    atkType string
    power int
    accuracy int
    secondEffectRate float32
    secondEffect string
}

// maps move name to information
const moveList map[string][Move]{
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
    "Flash Cannon": Move{"Flash Cannon", "Steel", "spatk", 80, 100, 0.1, "spdef"}
}

type Pokemon struct {
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

const PokemonList map[string]Pokemon{ 
    "Spiritomb": Pokemon{"Spiritomb", 108, "Ghost", "Dark", 50, 92, 108, 92, 108, 35, [4]string{"Dark Pulse", "Shadow Ball", "Bug Buzz", "Psychic"}}, 
    "Lucario": Pokemon{"Lucario", 116, "Fighting", "Steel", 70, 110, 70, 115, 70, 90, [4]string{"Strength", "Aura Sphere", "Shadow Ball", "Stone Edge"}}, 
    "Togekiss": Pokemon{"Togekiss", 175, "Normal", "Flying", 85, 50, 95, 120, 115, 80, [4]string{"Air Slash", "Aura Sphere", "Surf", "Thunderbolt"}}, 
    "Roserade": Pokemon{"Roserade", 27, "Grass", "Poison", 60, 70, 55, 125, 105, 90, [4]string{"Energy Ball", "Sludge Bomb", "Extrasensory", "Shadow Ball"}}, 
    "Milotic": Pokemon{"Milotic", 139, "Water", nil, 95, 60, 79, 100, 125, 81, [4]string{"Surf", "Ice Beam", "Dragon Pulse", "Flamethrower"}}, 
    "Garchomp": Pokemon{"Garchomp", 111, "Dragon", "Ground", 108, 130, 95, 80, 85, 102, [4]string{"Dragon Claw", "Earthquake", "Flamethrower", "Rock Slide"}},
    "Torterra": Pokemon{"Torterra", 3, "Grass", "Ground", 95, 109, 105, 75, 85, 56, [4]string{"Energy Ball", "Earthquake", "Crunch", "Rock Slide"}}, 
    "Infernape": Pokemon{"Infernape", 6, "Fire", "Fighting", 76, 104, 71, 104, 71, 108, [4]string{"Flamethrower", "Brick Break", "Aura Sphere", "Poison Jab"}}, 
    "Empoleon": Pokemon{"Empoleon", 9, "Water", "Steel", 84, 86, 88, 111, 101, 60, [4]string{"Surf", "Flash Cannon", "Aerial Ace", "Strength"}}, 
    "Staraptor": Pokemon{"Staraptor", 12, "Normal", "Flying", 85, 120, 70, 50, 50, 100, [4]string{"Strength", "Aerial Ace", "Brick Break", "Leaf Blade"}}, 
    "Bibarel": Pokemon{"Bibarel", 14, "Normal", "Water", 79, 85, 60, 55, 60, 71, [4]string{"Strength", "Waterfall", "Aerial Ace", "Brick Break"}}, 
    "Luxray": Pokemon{"Luxray", 19, "Electric", nil, 80, 120, 79, 95, 79, 70, [4]string{"Spark", "Crunch", "Brick Break", "Strength"}}, 
    "Alakazam": Pokemon{"Alakazam", 22, "Psychic", nil, 55, 50, 45, 135, 85, 120, [4]string{"Psychic", "Shadow Ball", "Dark Pulse", "Aura Sphere"}}, 
    "Gyarados": Pokemon{"Gyarados", 24, "Water", "Flying", 95, 125, 79, 60, 100, 81, [4]string{"Waterfall", "Aerial Ace", "Earthquake", "Hyper Voice"}}, 
    "Steelix": Pokemon{"Steelix", 35, "Steel", "Ground", 75, 85, 200, 55, 65, 30, [4]string{"Iron Head", "Earthquake", "Strength", "Stone Edge"}}, 
    "Machamp": Pokemon{"Machamp", 42, "Fighting", nil, 90, 130, 80, 65, 85, 55, [4]string{"Brick Break", "Rock Slide", "Strength", "Earthquake"}}, 
    "Gastrodon": Pokemon{"Gastrodon", 61, "Water", "Ground", 111, 83, 68, 92, 82, 39, [4]string{"Earth Power", "Surf", "Ice Beam", "Sludge Bomb"}}, 
    "Heracross": Pokemon{"Heracross", 62, "Bug", "Fighting", 80, 125, 75, 40, 95, 85, [4]string{"Brick Break", "X-scissor", "Earthquake", "Night Slash"}}, 
    "Gengar": Pokemon{"Gengar", 71, "Ghost", "Poison", 60, 65, 60, 130, 75, 110, [4]string{"Shadow Ball", "Sludge Bomb", "Dark Pulse", "Aura Sphere"}}, 
    "Honchkrow": Pokemon{"Honchkrow", 75, "Dark", "Flying", 100, 125, 52, 105, 52, 71, [4]string{"Night Slash", "Aerial Ace", "Shadow Ball", "Dark Pulse"}}, 
    "Rapidash": Pokemon{"Rapidash", 91, "Fire", nil, 65, 100, 70, 80, 80, 105, [4]string{"Flame Wheel", "Poison Jab", "Strength", "Iron Head"}}, 
    "Pikachu": Pokemon{"Pikachu", 104, "Electric", nil, 35, 55, 50, 50, 40, 90, [4]string{"Thunderbolt", "Flamethrower", "Surf", "Aura Sphere"}}, 
    "Snorlax": Pokemon{"Snorlax", 113, "Normal", nil, 160, 110, 65, 65, 110, 30, [4]string{"Strength", "Rock Slide", "Crunch", "Earthquake"}}, 
    "Drapion": Pokemon{"Drapion", 128, "Poison", "Dark", 70, 90, 110, 60, 75, 95, [4]string{"Poison Jab", "Night Slash", "X-scissor", "Earthquake"}}, 
    "Toxicroak": Pokemon{"Toxicroak", 130, "Poison", "Fighting", 83, 106, 65, 86, 65, 85, [4]string{"Poison Jab", "Brick Break", "Extrasensory", "Night Slash"}}, 
    "Abomasnow": Pokemon{"Abomasnow", 143, "Grass", "Ice", 90, 92, 75, 92, 85, 60, [4]string{"Avalanche", "Seed Bomb", "Ice Beam", "Earthquake"}}, 
    "Weavile": Pokemon{"Weavile", 145, "Dark", "Ice", 70, 120, 65, 45, 85, 125, [4]string{"Night Slash", "Avalanche", "X-scissor", "Iron Head"}}, 
    "Gardevoir": Pokemon{"Gardevoir", 159, "Psychic", nil, 68, 65, 65, 125, 115, 80, [4]string{"Psychic", "Shadow Ball", "Aura Sphere", "Energy Ball"}}, 
    "Gallade": Pokemon{"Gallade", 160, "Psychic", "Fighting", 68, 125, 65, 65, 115, 80, [4]string{"Leaf Blade", "Psycho Cut", "Brick Break", "Strength"}}, 
    "Altaria": Pokemon{"Altaria", 172, "Dragon", "Flying", 75, 70, 90, 70, 105, 80, [4]string{"Dragon Pulse", "Dragon Claw", "Aerial Ace", "Ice Beam"}}, 
    "Houndoom": Pokemon{"Houndoom", 177, "Dark", "Fire", 75, 90, 50, 110, 80, 95, [4]string{"Dark Pulse", "Flamethrower", "Energy Ball", "Earth Power"}}, 
    "Magnezone": Pokemon{"Magnezone", 180, "Electric", "Steel", 70, 70, 115, 130, 90, 60, [4]string{"Flash Cannon", "Thunderbolt", "Hyper Voice", "Bug Buzz"}}, 
    "Rhyperior": Pokemon{"Rhyperior", 188, "Ground", "Rock", 115, 140, 130, 55, 55, 40, [4]string{"Earthquake", "Stone Edge", "Poison Jab", "Brick Break"}}, 
    "Dusknoir": Pokemon{"Dusknoir", 191, "Ghost", nil, 45, 100, 135, 65, 135, 45, [4]string{"Shadow Punch", "Night Slash", "Rock Slide", "Strength"}}, 
    "Porygon-z": Pokemon{"Porygon-z", 194, "Normal", nil, 85, 80, 70, 135, 75, 90, [4]string{"Hyper Voice", "Thunderbolt", "Dark Pulse", "Ice Beam"}}, 
    "Scizor": Pokemon{"Scizor", 196, "Bug", "Steel", 70, 130, 100, 55, 80, 65, [4]string{"Iron Head", "X-scissor", "Strength", "Night Slash"}}, 
    "Electivire": Pokemon{"Electivire", 199, "Electric", nil, 75, 123, 67, 95, 85, 95, [4]string{"Spark", "Earthquake", "Iron Head", "Rock Slide"}}, 
    "Magmortar": Pokemon{"Magmortar", 202, "Fire", nil, 75, 95, 67, 125, 95, 83, [4]string{"Flamethrower", "Dark Pulse", "Thunderbolt", "Psychic"}}, 
    "Mamoswine": Pokemon{"Mamoswine", 205, "Ice", "Ground", 110, 130, 80, 70, 60, 80, [4]string{"Avalanche", "Earthquake", "Strength", "Stone Edge"}}, 
    "Froslass": Pokemon{"Froslass", 208, "Ice", "Ghost", 70, 80, 70, 80, 70, 110, [4]string{"Shadow Ball", "Ice Beam", "Thunderbolt", "Psychic"}}
} 

const winMessages []string{"That was excellent. Truly, an outstanding battle. You gave the support your Pokémon needed to maximize their power. And you guided them with certainty to secure victory. You have both passion and calculating coolness. Together, you and your Pokémon can overcome any challenge that may come your way. Those are the impressions I got from our battle. I'm glad I got to take part in the crowning of Sinnoh's new Champion! Come with me. We'll take the lift."}   // win messages
const loseMessages []string{"Smell ya later!", "Better luck next time", "Keep training", "Time to soft-reset", "You whited out...", "Come back when you're stronger", "Do or do not...there is no try", "Looks like you're blasting off again"}   // lose messages

type UserInput struct {
    action string
    pokemon &Pokemon
    move string
}