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
		cellGrid := cells[(game.Cols * row64):(game.Cols * (row64 + 1))]
		game.Grid[row] = cellGrid
	}
}

func setCellValues(game *domain.Game) {
	for i, column := range game.Grid {
		for j, cell := range column {
			game.Grid[i][j].Column = int64(i)
			game.Grid[i][j].Row = int64(j)
			if cell.Mine {
				setAdjacentValues(game, cell)
			}
		}
	}
}

func setAdjacentValues(game *domain.Game, cell domain.Cell) {
	for z := cell.Row - 1; z < cell.Row+2; z++ {
		if z < 0 || int64(z) > game.Rows-1 {
			continue
		}
		for w := cell.Column - 1; w < cell.Column+2; w++ {

			if w < 0 || int64(w) > game.Cols-1 {
				continue
			}
			if z == cell.Row && w == cell.Column {
				continue
			}
			game.Grid[z][w].StartingAdjacentBombs++
		}
	}
}
