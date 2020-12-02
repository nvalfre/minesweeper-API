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
	numCells := game.Cols * game.Rows
	cells := make(domain.CellGrid, numCells)

	generateRandomPositions(game, numCells, cells)
	setCellValues(game)
}

func generateRandomPositions(game *domain.Game, numCells int64, cells domain.CellGrid) {
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
		game.Grid[row] = cells[(game.Cols * row64):(game.Cols * (row64 + 1))]
	}
}

func setCellValues(game *domain.Game) {
	for i, row := range game.Grid {
		for j, cell := range row {
			if cell.Mine {
				setAdjacentValues(game, i, j)
			}
		}
	}
}

func setAdjacentValues(game *domain.Game, i, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || int64(z) > game.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {

			if w < 0 || int64(w) > game.Cols-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			game.Grid[z][w].Value++
		}
	}
}
