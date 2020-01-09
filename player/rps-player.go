package rps_player

import (
	"fmt"
	"math/rand"
	"rock-paper-scisors/rps"
	"strings"
)

type RPSPlayer struct {
	game rps.RPSGame
}

func CreateNewGame() *RPSPlayer {
	return &RPSPlayer{}
}


var (
	choices                      = []string{"rock", "paper", "scissors"}
	running                      bool
	player1choice, player2choice string
)

func randomAnswer() string {
	return choices[rand.Intn(len(choices))]
}

func isValid(choice string) bool {
	for _, c := range choices {
		if choice == c {
			return true
		}
	}
	return false
}

func choiceCompare(choice1, choice2 string) string {
	switch {
	case !isValid(choice1):
		return "Player one's choice was not valid!"
	case !isValid(choice2):
		return "Player two's choice was not valid!"
	case choice1 == choice2:
		return "There was a tie!"
	case (choice1 == "rock" && choice2 == "scissors") || (choice1 == "paper" && choice2 == "rock") || (choice1 == "scissors" && choice2 == "paper"):
		return "Player one wins!"
	default:
		return "Player two wins!"
	}
}
func playAgain() bool {
	fmt.Println("Would you like to play again? (y/n)")

	var again string
	fmt.Scanf("%s", &again)
	again = strings.ToLower(again)

	if again == "y" {
		return true
	} else if again == "n" {
		return false
	} else {
		fmt.Println("That's an invalid response!")
		return playAgain()
	}
}

func (player *RPSPlayer) Play() {
	running = true
	for running == true {
		fmt.Println("Would you like to play with 0, 1, or 2 players?")

		var players string
		fmt.Scanf("%s", &players)
		if players == "1" || players == "2" {
			fmt.Println("What is your choice? (rock, paper, or scissors)")
			fmt.Scanf("%s", &player1choice)
		} else {
			fmt.Println("The player 1 computer is choosing...")
			player1choice = randomAnswer()
		}
		if players == "2" {
			fmt.Println("what is your opponents choice?")
			fmt.Scanf("%s", &player2choice)
		} else {
			fmt.Println("The player 2 computer is choosing...")
			player2choice = randomAnswer()
		}
		player1choice = strings.ToLower(player1choice)
		player2choice = strings.ToLower(player2choice)

		fmt.Println(choiceCompare(player1choice, player2choice))
		fmt.Println("Player 1 chose ", player1choice)
		fmt.Println("Player 2 chose ", player2choice)

		running = playAgain()
	}
}


