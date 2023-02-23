package main

import (
	"fmt"
	"math/rand"
	"time"
)

func play_game(switchDoor bool) bool {
	// Initialize the random number generator with the current time
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	doors := []int{1, 2, 3}
	prize := doors[rng.Intn(3)]
	chosen := 0

	// Create a new source of random
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select a random door
	input := rng.Intn(3)
	chosen = input

	// Open one of the other two doors to reveal a goat
	var opened int
	for _, door := range doors {
		if door != chosen && door != prize {
			opened = door
			break
		}
	}

	// Switch the door if switchDoor equals true
	if switchDoor {
		for _, door := range doors {
			if door != chosen && door != opened {
				chosen = door
				break
			}
		}
	}

	// Return if the prize is won
	return chosen == prize
}

func main() {
	// How many games to play
	gamesToPlay := 10000
	games := make([]int, gamesToPlay)

	// Should the door be switched?
	switchDoor := true

	for i := 0; i < gamesToPlay; i++ {
		won := play_game(switchDoor)
		if won {
			games[i] = 1
		} else {
			games[i] = 0
		}
	}

	totalWon := 0
	totalLost := 0

	// Count the number of won and lost games
	for _, num := range games {
		if num == 0 {
			totalLost++
		} else {
			totalWon++
		}
	}

	switchedString := "are"
	if !switchDoor {
		switchedString = "isn't"
	}

	// Calculate the winning chance
	winningChance := float64(totalWon) / float64(gamesToPlay) * 100.0
	fmt.Printf("In %d games\nGames won: %d\nGames Lost: %d\n", gamesToPlay, totalWon, totalLost)
	fmt.Printf("This gives you a winning chance of %.2f%% percent if the doors %s switched\n", winningChance, switchedString)
}
