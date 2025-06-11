//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
)

// builtins are a good thing

var p = fmt.Println

type Player struct {
	name                                 string
	health, maxHealth, energy, maxEnergy int
}

func (player *Player) heal(healBy int) {
	player.health = min(player.health+healBy, player.maxHealth)
}

func (player *Player) damage(hurtBy int) {
	player.health = max(player.health-hurtBy, 0)
}

func (player *Player) powerUp(powerUpBy int) {
	player.energy = min(player.energy+powerUpBy, player.maxEnergy)
}

func (player *Player) blast(consume int) {
	player.energy = max(player.energy-consume, 0)
}

func main() {
	var playerOne = Player{name: "Timothy", health: 100, maxHealth: 1000, energy: 100, maxEnergy: 1000}
	p(playerOne)

	playerOne.heal(100)
	playerOne.powerUp(100)

	p(playerOne)

	playerOne.damage(199)
	p(playerOne)

	playerOne.heal(10000)
	playerOne.blast(9999)
	p(playerOne)

}
