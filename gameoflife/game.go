package gameoflife

import "time"

// Game is the point of entry of the whole game
type Game struct {
	*board
}

// Play is the entry point to start a new game
func (g *Game) Play(boardWidth int, boardHeight int, patternFilePath string, neighborStrategy string) {
	parser := patternParser{patternFilePath: patternFilePath}

	g.board = newBoard(boardWidth, boardHeight, parser.extractCoordinates(), neighborStrategy)

	g.start()
}

func (g *Game) start() {
	for {
		g.Display()
		time.Sleep(50 * time.Millisecond)
		g.Step()
	}
}
