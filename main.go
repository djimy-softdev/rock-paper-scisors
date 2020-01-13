package main

import (
	"fmt"
	"math/rand"
	rps_player "rock-paper-scisors/player"
	"rock-paper-scisors/rps"
	"time"
)

func main() {
	player := rps_player.CreateNewGame(&rps.RPSGameImp{})


	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to Rock Paper Scissors!")

	player.Play()

	fmt.Println("\nThanks for playing Rock Paper Scissors.")
}