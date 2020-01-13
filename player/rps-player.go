package rps_player

import (
	"fmt"
	"math/rand"
	"rock-paper-scisors/rps"
	"strconv"
	"strings"
)

type RPSPlayer struct {
	game rps.RPSGame
}

/*
 * Params: an implementation of a "rock paper scissor" (rps) game
 * Returns a rps game player object
 */
func CreateNewGame(game rps.RPSGame) *RPSPlayer {
	return &RPSPlayer{game}
}

func playAgain() (bool, error) {
	fmt.Print("\nWould you like to play again? (y/n)")

	var again string
	_, err := fmt.Scanf("%s", &again)
	if err != nil {
		return false, err
	}

	again = strings.ToLower(again)

	if again == "y" {
		return true, nil
	}

	return false, nil
}

func getComputerChoice(name string, validator func(choice rps.RPSChoice)bool) rps.RPSChoice {
	choice := rps.RPSChoice(rand.Uint32()%3)
	fmt.Printf("%v's choice: %v\n", name, choice)
	return choice
}

func getPlayerChoice(name string, validator func(choice rps.RPSChoice)bool) rps.RPSChoice {
	var playerChoice string
	var choice rps.RPSChoice

	for true {
		fmt.Printf("%v's choice: ", name)
		_, err := fmt.Scanf("%s", &playerChoice)
		if err != nil {
			fmt.Println("Wrong input!")
			continue
		}

		a, err := strconv.Atoi(playerChoice)
		if choice = rps.RPSChoice(a); err != nil || !validator(choice) {
			fmt.Println("Wrong input!")
			continue
		}

		break
	}

	return choice
}

func getChoices(players int, playerNames []string, validator func(choice rps.RPSChoice)bool) ([]rps.RPSChoice, error) {
	fmt.Print("\nChoices: rock[1], paper[2], scissors[3] \n")
	choicesFunc := []func(string, func(choice rps.RPSChoice)bool)rps.RPSChoice{getComputerChoice, getComputerChoice}
	for i := 0; i < players; i++ {
		choicesFunc[i] = getPlayerChoice
	}

	choicesVal := make([]rps.RPSChoice, len(choicesFunc))
	for index, choice := range choicesFunc {
		val := choice(playerNames[index], validator)
		choicesVal[index] = val
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
		fmt.Println("It's a draw.")
		return
	}

	if verdict == rps.Win {
		fmt.Printf("%v wins.\n", playerNames[0])
		return
	}

	fmt.Printf("%v wins. \n", playerNames[1])
}

/*
 * The entry point of the game, this function will start the game
 */
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
		choices, err := getChoices(players, playerNames, player.game.ValidateChoice)

		for err != nil {
			fmt.Printf("Wrong input!\n")
			break
		}

		verdict, _ := player.game.GetVerdict(choices[0], choices[1])
		displayVerdict(playerNames, verdict)

		keepOn, err := playAgain()
		if err != nil {
			fmt.Println("Fail to read input!")
		}

		if !keepOn {
			fmt.Println("Goodbye!!!")
			break
		}
	}
}


