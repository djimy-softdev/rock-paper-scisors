package rps

import (
	"errors"
	"testing"
)

func TestRps(t *testing.T) {
	testCases := []struct{
		player1Choice RPSChoice
		player2Choice RPSChoice
		verdict RPSVerdict
		err error
	}{
		{rock, rock, Draw, nil},
		{rock, paper, Lose, nil},
		{rock, scisor, Win, nil},

		{paper, rock, Win, nil},
		{paper, paper, Draw, nil},
		{paper, scisor, Lose,nil},

		{scisor, rock, Lose, nil},
		{scisor, paper, Win, nil},
		{scisor, scisor, Draw, nil},

		{4, 3, Lose, errors.New("WRONG ARGUMENT")},
	}

	game := &RPSGameImp{}
	for _, tc := range testCases {
		testVerdict, err := game.GetVerdict(tc.player1Choice, tc.player2Choice)

		if err != tc.err  && testVerdict != tc.verdict {
			t.Errorf("input %v %v --- output: %v --- expected: %v \n", tc.player1Choice, tc.player2Choice, testVerdict, tc.verdict)
		}
	}
}