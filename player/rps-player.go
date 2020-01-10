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

func CreateNewGame(game rps.RPSGame) *RPSPlayer {
	return &RPSPlayer{game}
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

func getComputerChoice(index int) (rps.RPSChoice, error) {
	choice := rps.RPSChoice(rand.Uint32()%3)
	fmt.Printf("Computer%v choice: %v \n", index, choice)
	return choice, nil
}

func getPlayerChoice(index int) (rps.RPSChoice, error) {
	var playerChoice rps.RPSChoice
	fmt.Printf("\nPlayer%v choice: ", index)
	_, err := fmt.Scanf("%s", &playerChoice)
	return playerChoice, err
}

func getChoices(players int) ([]rps.RPSChoice, error) {
	fmt.Print("\nChoices: rock[1], paper[2], scissors[3] \n")
	choicesFunc := []func(int)(rps.RPSChoice, error){getComputerChoice, getComputerChoice}
	for i := 0; i < players; i++ {
		choicesFunc[i] = getPlayerChoice
	}

	choicesVal := make([]rps.RPSChoice, len(choicesFunc))
	var err error
	for index, choice := range choicesFunc {
		choicesVal[index], err = choice(index)
		if err != nil {
			return nil, err
		}
	}

	return choicesVal, nil
}

func generatePlayerNames(numPlayers int) []string{
	var playerNames []string
	switch numPlayers {
	case 0:
		playerNames = []string{"computer1", "computer2"}
	case 1:
		playerNames = []string{"player", "computer"}
	default:
		playerNames = []string{"player1", "player2"}
	}

	return playerNames
}

func displayVerdict(playerNames []string, verdict rps.RPSVerdict) {
	if verdict == rps.Draw {
		fmt.Println("\nIt's a draw. \n")
		return
	}

	if verdict == rps.Win {
		fmt.Printf("%v wins.\n", playerNames[0])
		return
	}

	fmt.Printf("%v wins. \n", playerNames[1])
}

func (player *RPSPlayer) Play() {

	fmt.Println("Would you like to play with 0, 1, or 2 players?")
	var players int
	_, err := fmt.Scanf("%d", &players)
	if err != nil {
		fmt.Println("Fail to ")
	}

	if players > 2 || players < 0 {
		fmt.Println("Invalid argument.")
		return
	}

	for true {
		playerNames := generatePlayerNames(players)
		choices, _ := getChoices(players)

		verdict, _ := player.game.GetVerdict(choices[0], choices[1])
		displayVerdict(playerNames, verdict)

		if !playAgain() {
			fmt.Println("Goodbye!!!")
			break
		}
	}
}


