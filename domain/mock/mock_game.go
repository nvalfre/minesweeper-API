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
	MockCell01Bomb = domain.Cell{
		Row:    0,
		Column: 1,
		Mine:   true,
	}
	MockCell12Bomb = domain.Cell{
		Row:    1,
		Column: 2,
		Mine:   true,
	}
	MockCell21bomb = domain.Cell{
		Row:    2,
		Column: 1,
		Mine:   true,
	}
	MockCell22   = getMockCell(2, 2)
	Grid         = []domain.CellGrid{{MockCell0, MockCell01}, {MockCell11, MockCell12}, {MockCell21, MockCell22}}
	GridWithBomb = []domain.CellGrid{{MockCell0, MockCell01}, {MockCell11, MockCell12}, {MockCell21bomb, MockCell22}}

	Game = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       Grid,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
	GameBomb = domain.Game{
		Rows:       3,
		Cols:       2,
		Mines:      2,
		Grid:       GridWithBomb,
		GameStatus: game_status.GameStatus{Status: game_status.Started},
	}
)

func getMockCell(row, column int64) domain.Cell {
	return domain.Cell{
		Row:    row,
		Column: column,
	}
}
