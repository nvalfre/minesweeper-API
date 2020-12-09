package mock

import (
	"minesweeper-API/domain"
	"minesweeper-API/domain/game_status"
)

var (
	MockCell0      = getMockCell(0, 0)
	MockCell01     = getMockCell(0, 1)
	MockCell11     = getMockCell(1, 1)
	MockCell12     = getMockCell(1, 2)
	MockCell21     = getMockCell(2, 1)
	MockCell21bomb = domain.Cell{
		Row:    0,
		Column: 0,
		Mine:   false,
	}
	MockCell22   = getMockCell(2, 2)
	Grid         = initGridDefault()
	GridWithBomb = initGridBomb()

	Game = domain.Game{
		Rows:       2,
		Cols:       2,
		Mines:      2,
		Grid:       Grid,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	GameBomb = domain.Game{
		Rows:       2,
		Cols:       2,
		Mines:      2,
		Grid:       GridWithBomb,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
)

func initGridBomb() []domain.CellGrid {
	return []domain.CellGrid{{MockCell0, MockCell01}, {MockCell11, MockCell12}, {MockCell21bomb, MockCell22}}
}

func initGridDefault() []domain.CellGrid {
	return []domain.CellGrid{{MockCell0, MockCell01}, {MockCell11, MockCell12}, {MockCell21, MockCell22}}
}

func getMockCell(row, column int64) domain.Cell {
	return domain.Cell{
		Row:    row,
		Column: column,
	}
}
