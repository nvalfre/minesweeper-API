package services

import (
	"math/rand"
	"minesweeper-API/domain"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func BuildBoard(game *domain.Game) {
	generateRandomPositions(game)
	setCellValues(game)
}

func generateRandomPositions(game *domain.Game) {
	numCells := game.Rows * game.Cols

	cells := make(domain.CellGrid, numCells)

	i := 0
	for int64(i) < game.Mines {
		idx := rand.Intn(int(numCells))
		if !cells[idx].Mine {
			cells[idx].Mine = true
			i++
		}
	}

	game.Grid = make([]domain.CellGrid, game.Rows)
	for row := range game.Grid {
		row64 := int64(row)
		cellGrid := cells[(game.Cols * row64):(game.Cols * (row64 + 1))]
		game.Grid[row] = cellGrid
	}
}

func setCellValues(game *domain.Game) {
	for i, row := range game.Grid {
		for j, _ := range row {
			game.Grid[i][j].Column = int64(j)
			game.Grid[i][j].Row = int64(i)
		}
	}
}
