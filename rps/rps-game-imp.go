package rps

import "errors"

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

func (d *RPSGameImp) GetVerdict(player1Choice RPSChoice, player2Choice RPSChoice) (error, RPSVerdict) {
	if player1Choice >= RPSChoice(len(choices)) || player2Choice >= RPSChoice(len(choices)) {
		return errors.New("WRONG ARGUMENT"), 0
	}

	verdict =

	return Draw
}