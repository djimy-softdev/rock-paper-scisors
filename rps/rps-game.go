package rps

type RPSGame interface {
	GetVerdict(player1 RPSChoice, player2 RPSChoice) (RPSVerdict, error)
}

type RPSVerdict uint32

type RPSChoice uint32
