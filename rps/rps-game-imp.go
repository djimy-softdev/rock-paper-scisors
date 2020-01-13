package rps

import (
	"errors"
)

type RPSGameImp struct {
}

const (
	Lose RPSVerdict = 0
	Draw RPSVerdict = 1
	Win  RPSVerdict = 2

	rock RPSChoice = 0
	paper RPSChoice = 1
	scisor  RPSChoice = 2
)

var verdicts = []RPSVerdict{
	Lose,
	Draw,
	Win,
}

var choices = []RPSChoice{
	rock,
	paper,
	scisor,
}

/*
 * Returns the verdict of the game for player1
 * Possible verdicts: Win, Lose, Draw
 */
func (d *RPSGameImp) GetVerdict(player1Choice RPSChoice, player2Choice RPSChoice) (RPSVerdict, error) {
	if player1Choice >= RPSChoice(len(choices)) || player2Choice >= RPSChoice(len(choices)) {
		return Lose, errors.New("WRONG ARGUMENT")
	}

	if player1Choice == player2Choice {
		return Draw, nil
	}

	if (player1Choice + 1) % 3 == player2Choice {
		return Lose, nil
	}

	return Win, nil
}

/*
 * Returns whether or not the player's choice is valid
 * Possible values: true, false
 */
func (d *RPSGameImp) ValidateChoice(playerChoice RPSChoice) bool {
	if int(playerChoice) >= len(choices) {
		return false
	}

	return true
}