package rps

type RPSGame interface {
	GetVerdict(player1 RPSChoice, player2 RPSChoice) (RPSVerdict, error)
	ValidateChoice(playerChoice RPSChoice) bool
}

type RPSVerdict uint32

type RPSChoice uint32
