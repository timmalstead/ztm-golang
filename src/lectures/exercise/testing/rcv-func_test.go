//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
//

package main

import "testing"

var player = Player{name: "Timothy", health: 100, maxHealth: 1000, energy: 100, maxEnergy: 1000}

func TestMin(t *testing.T) {
	player.blast(10000)
	player.damage(10000)

	if player.health != 0 {
		t.Errorf("health should be 0 instead of %v", player.health)
	}

	if player.energy != 0 {
		t.Errorf("energy should be 0 instead of %v", player.energy)
	}
}

func TestMax(t *testing.T) {
	player.powerUp(10000)
	player.heal(10000)

	if player.health > 1000 {
		t.Errorf("health: %v, should not be greater than maxHealth: %v", player.health, player.maxHealth)
	}

	if player.energy > 1000 {
		t.Errorf("energy: %v, should not be greater than maxEnergy: %v", player.energy, player.maxEnergy)
	}
}
