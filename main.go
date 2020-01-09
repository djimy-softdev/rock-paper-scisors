package main

import (
	"fmt"
	"math/rand"
	rps_player "rock-paper-scisors/player"
	"time"
)

func main() {
	player := rps_player.CreateNewGame()


	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to Rock Paper Scissors!")

	player.Play()

	fmt.Println("Thanks for playing Rock Paper Scissors, by ParrotCaws")
	fmt.Println("Fork me on github! github.com/ParrotCaws")
}